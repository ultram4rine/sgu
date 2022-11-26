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
	"github.com/diskfs/go-diskfs/partition/gpt"
	"golang.org/x/sys/windows"
)

var CNT = 1

func getEntryOffset(f *pe.File, name string) (offset uint64, err error) {
	for _, s := range f.Symbols {
		if s.Name == name {
			sect := f.Sections[s.SectionNumber-1]
			return uint64(sect.Offset + s.Value), nil
		}
	}
	return 0, fmt.Errorf("can't find symbol '%s'", name)
}

func patchPE(peName string, offset uint64) ([]byte, error) {
	f, err := os.Open(peName)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(data[offset:], 2)

	return data, nil
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

func main() {
	dev, err := determineDevice()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(dev /* , os.O_WRONLY|os.O_APPEND, 0 */)
	if err != nil {
		log.Fatalf("can't open %s: %v", dev, err)
	}

	tbl, err := partition.Read(f, 512, 512)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tbl.(*gpt.Table).GUID)

	parts := tbl.GetPartitions()
	buf := new(bytes.Buffer)
	_, err = parts[len(parts)-1].ReadContents(f, buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf.Bytes())
	bt := buf.Bytes()
	bt[0] = 1
	buf.Truncate(0)
	buf.Write(bt)
	log.Println(buf.Bytes())

	if CNT == 1 {
		fmt.Println("first run")
	} else if CNT == 2 {
		fmt.Println("second run")
	}

	var exeName = os.Args[0]

	self, err := pe.Open(exeName)
	if err != nil {
		log.Fatal(err)
	}

	offset, err := getEntryOffset(self, "main.CNT")
	if err != nil {
		log.Fatalf("can't find counter object in ELF file: %v", err)
	}

	patchedData, err := patchPE(exeName, offset)
	if err != nil {
		log.Fatalf("can't patch PE file: %v", err)
	}

	tmpExeName := exeName + ".tmp"
	newSelf, err := os.Create(tmpExeName)
	if err != nil {
		log.Fatalf("can't create temporary file: %v", err)
	}

	if _, err := newSelf.Write(patchedData); err != nil {
		log.Fatalf("can't write patched PE: %v", err)
	}

	cmd := exec.Command("cmd.exe", "/C", fmt.Sprintf("start cmd /C move /y %s %s && exit", tmpExeName, exeName))
	log.Println(cmd.Start())
}
