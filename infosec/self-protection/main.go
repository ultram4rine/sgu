package main

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/diskfs/go-diskfs/partition"
	"golang.org/x/sys/windows"
)

var CNT = 1

func main() {
	dev, err := determineDevice()
	if err != nil {
		log.Fatal(err)
	}

	if CNT == 1 {
		fmt.Println("first run")

		if err := markDevice(dev); err != nil {
			log.Fatal(err)
		}

		var exeName = os.Args[0]
		if err := patchPE(exeName); err != nil {
			log.Fatalf("can't patch PE file: %v", err)
		}
	} else if CNT == 2 {
		ok, err := checkDevice(dev)
		if err != nil {
			log.Fatal("failed to check device: %v", err)
		}
		if !ok {
			log.Fatal("wrong device")
		}
		fmt.Println("second run")
	} else {
		fmt.Println("unreachable block")
	}
}

func determineDevice() (string, error) {
	sysRoot := os.Getenv("SYSTEMROOT")
	rootDrive := sysRoot[:2]

	mode := uint32(windows.FILE_SHARE_READ | windows.FILE_SHARE_WRITE | windows.FILE_SHARE_DELETE)
	flags := uint32(windows.FILE_ATTRIBUTE_READONLY)
	f, err := windows.CreateFile(windows.StringToUTF16Ptr("\\\\.\\"+rootDrive), windows.GENERIC_READ, mode, nil, windows.OPEN_EXISTING, flags, 0)
	if err != nil {
		return "", err
	}

	controlCode := uint32(5636096) // IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS
	size := uint32(16 * 1024)
	vols := make([]byte, size)
	var bytesReturned uint32
	if err := windows.DeviceIoControl(f, controlCode, nil, 0, &vols[0], size, &bytesReturned, nil); err != nil {
		return "", err
	}
	if uint(binary.LittleEndian.Uint32(vols)) != 1 {
		return "", fmt.Errorf("could not identify physical drive for %s", rootDrive)
	}

	diskId := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(vols[8:])), 10)

	drive := "\\\\.\\PhysicalDrive" + diskId
	f, err = windows.CreateFile(windows.StringToUTF16Ptr(drive), windows.GENERIC_READ, mode, nil, windows.OPEN_EXISTING, flags, 0)
	if err != nil {
		log.Fatal(err)
	}

	mbr := make([]byte, 512)
	if err = windows.ReadFile(f, mbr, &bytesReturned, nil); err != nil {
		log.Fatal(err)
	}

	if mbr[510] != 0x55 || mbr[511] != 0xaa {
		return "", fmt.Errorf("MBR not detected on physical drive for %s (%s)", rootDrive, drive)
	}

	return drive, nil
}

func checkDevice(dev string) (bool, error) {
	f, err := os.OpenFile(dev, os.O_RDWR, 0600)
	if err != nil {
		return false, fmt.Errorf("can't open %s: %v", dev, err)
	}
	defer f.Close()

	tbl, err := partition.Read(f, 512, 512)
	if err != nil {
		return false, err
	}

	parts := tbl.GetPartitions()
	part := parts[len(parts)-1]
	buf := new(bytes.Buffer)

	if _, err = part.ReadContents(f, buf); err != nil {
		return false, err
	}

	bt := buf.Bytes()
	// First 4 bytes of Partition entry №1 must be 1 2 3 4.
	if bt[494] == 1 && bt[495] == 2 && bt[496] == 3 && bt[497] == 4 {
		return true, nil
	} else {
		return false, nil
	}
}

func markDevice(dev string) error {
	f, err := os.OpenFile(dev, os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("can't open %s: %v", dev, err)
	}
	defer f.Close()

	tbl, err := partition.Read(f, 512, 512)
	if err != nil {
		return err
	}

	parts := tbl.GetPartitions()
	part := parts[len(parts)-1]
	buf := new(bytes.Buffer)
	if _, err = part.ReadContents(f, buf); err != nil {
		return err
	}

	bt := buf.Bytes()
	// Let's set first 4 bytes of Partition entry №1 to 1 2 3 4.
	bt[494] = 1
	bt[495] = 2
	bt[496] = 3
	bt[497] = 4
	buf.Truncate(0)
	buf.Write(bt)

	if _, err := part.WriteContents(f, buf); err != nil {
		return err
	}

	return nil
}

func getEntryOffset(f *pe.File, name string) (offset uint64, err error) {
	for _, s := range f.Symbols {
		if s.Name == name {
			sect := f.Sections[s.SectionNumber-1]
			return uint64(sect.Offset + s.Value), nil
		}
	}
	return 0, fmt.Errorf("can't find symbol '%s'", name)
}

func patchPE(peName string) error {
	self, err := pe.Open(peName)
	if err != nil {
		return err
	}

	offset, err := getEntryOffset(self, "main.CNT")
	if err != nil {
		return fmt.Errorf("can't find counter object in PE file: %v", err)
	}

	f, err := os.Open(peName)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	binary.LittleEndian.PutUint64(data[offset:], 2)

	tmpName := peName + ".tmp"
	newSelf, err := os.Create(tmpName)
	if err != nil {
		return fmt.Errorf("can't create temporary file: %v", err)
	}

	if _, err := newSelf.Write(data); err != nil {
		return fmt.Errorf("can't write patched PE: %v", err)
	}

	cmd := exec.Command("cmd.exe", "/C", fmt.Sprintf("start cmd /C move /y %s %s && exit", tmpName, peName))
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to move binary: %v", err)
	}

	return nil
}
