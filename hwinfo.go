package main

import (
	"errors"
	"net"
	"os/exec"
	"regexp"
	"strings"
)

//GetSysUUID 获取系统UUID
func GetSysUUID() (string, error) {
	return wmicGetID("csproduct", "UUID")
}

//GetBIOSID 获取BIOS ID
func GetBIOSID() (string, error) {
	return wmicGetID("bios", "SerialNumber")
}

//GetBaseBoardID 获取主板的id
func GetBaseBoardID() (string, error) {
	return wmicGetID("baseboard", "SerialNumber")
}

// GetCpuId 获取CPUiD
func GetCpuId() (string, error) {
	return wmicGetID("cpu", "ProcessorId")
}

func wmicGetID(a string, n string) (string, error) {
	cmd := exec.Command("wmic", a, "get", n)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	str := string(out)
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")

	if strings.Contains(str, n) {
		return str[len(n):], nil
	}
	return "", errors.New("id not found")
}

func GetMacAddr() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		return macAddr, nil
	}
	return "", errors.New("mac addr empty")
}
