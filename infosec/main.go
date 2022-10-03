package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/saferwall/pe"
	"golang.org/x/exp/slices"
)

var netFuncs = []string{
	"NetAlertRaise",
	"NetAlertRaiseEx",

	"NetApiBufferAllocate",
	"NetApiBufferFree",
	"NetApiBufferReallocate",
	"NetApiBufferSize",

	"NetFreeAadJoinInformation",
	"NetGetAadJoinInformation",

	"NetAddAlternateComputerName",
	"NetCreateProvisioningPackage",
	"NetEnumerateComputerNames",
	"NetGetJoinableOUs",
	"NetGetJoinInformation",
	"NetJoinDomain",
	"NetProvisionComputerAccount",
	"NetRemoveAlternateComputerName",
	"NetRenameMachineInDomain",
	"NetRequestOfflineDomainJoin",
	"NetRequestProvisioningPackageInstall",
	"NetSetPrimaryComputerName",
	"NetUnjoinDomain",
	"NetValidateName",

	"NetGetAnyDCName",
	"NetGetDCName",
	"NetGetDisplayInformationIndex",
	"NetQueryDisplayInformation",

	"NetGroupAdd",
	"NetGroupAddUser",
	"NetGroupDel",
	"NetGroupDelUser",
	"NetGroupEnum",
	"NetGroupGetInfo",
	"NetGroupGetUsers",
	"NetGroupSetInfo",
	"NetGroupSetUsers",

	"NetLocalGroupAdd",
	"NetLocalGroupAddMembers",
	"NetLocalGroupDel",
	"NetLocalGroupDelMembers",
	"NetLocalGroupEnum",
	"NetLocalGroupGetInfo",
	"NetLocalGroupGetMembers",
	"NetLocalGroupSetInfo",
	"NetLocalGroupSetMembers",

	"NetMessageBufferSend",
	"NetMessageNameAdd",
	"NetMessageNameDel",
	"NetMessageNameEnum",
	"NetMessageNameGetInfo",

	"NetFileClose",
	"NetFileEnum",
	"NetFileGetInfo",

	"NetRemoteComputerSupports",
	"NetRemoteTOD",

	"NetScheduleJobAdd",
	"NetScheduleJobDel",
	"NetScheduleJobEnum",
	"NetScheduleJobGetInfo",
	"GetNetScheduleAccountInformation",
	"SetNetScheduleAccountInformation",

	"NetServerDiskEnum",
	"NetServerEnum",
	"NetServerGetInfo",
	"NetServerSetInfo",

	"NetServerComputerNameAdd",
	"NetServerComputerNameDel",
	"NetServerTransportAdd",
	"NetServerTransportAddEx",
	"NetServerTransportDel",
	"NetServerTransportEnum",
	"NetWkstaTransportEnum",

	"NetUseAdd",
	"NetUseDel",
	"NetUseEnum",
	"NetUseGetInfo",

	"NetUserAdd",
	"NetUserChangePassword",
	"NetUserDel",
	"NetUserEnum",
	"NetUserGetGroups",
	"NetUserGetInfo",
	"NetUserGetLocalGroups",
	"NetUserSetGroups",
	"NetUserSetInfo",

	"NetUserModalsGet",
	"NetUserModalsSet",

	"NetValidatePasswordPolicyFree",
	"NetValidatePasswordPolicy",

	"NetWkstaGetInfo",
	"NetWkstaSetInfo",
	"NetWkstaUserEnum",
	"NetWkstaUserGetInfo",
	"NetWkstaUserSetInfo",

	"NetAccessAdd",
	"NetAccessCheck",
	"NetAccessDel",
	"NetAccessEnum",
	"NetAccessGetInfo",
	"NetAccessGetUserPerms",
	"NetAccessSetInfo",
	"NetAuditClear",
	"NetAuditRead",
	"NetAuditWrite",
	"NetConfigGet",
	"NetConfigGetAll",
	"NetConfigSet",
	"NetErrorLogClear",
	"NetErrorLogRead",
	"NetErrorLogWrite",
	"NetLocalGroupAddMember",
	"NetLocalGroupDelMember",
	"NetServiceControl",
	"NetServiceEnum",
	"NetServiceGetInfo",
	"NetServiceInstall",
	"NetWkstaTransportAdd",
	"NetWkstaTransportDel",
}

func main() {
	dir := ""
	if len(os.Args) < 2 {
		dir = "/mnt/c/Windows/System32"
	} else {
		dir = os.Args[1]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".exe") {
			continue
		}

		filename := filepath.Join(dir, file.Name())
		peFile, err := pe.New(filename, &pe.Options{})
		if err != nil {
			log.Fatalf("Error while opening file: %s, reason: %v", filename, err)
		}

		if err = peFile.Parse(); err != nil {
			log.Fatalf("Error while parsing file: %s, reason: %v", filename, err)
		}

	outer:
		for _, im := range peFile.Imports {
			for _, fun := range im.Functions {
				if slices.Contains(netFuncs, fun.Name) {
					fmt.Println(filename)
					break outer
				}
			}
		}
	}
}
