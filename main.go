package main

import "fmt"

func main() {
	mac, _ := GetMacAddr()
	fmt.Println("网卡MAC:", mac)
	board, _ := GetBaseBoardID()
	fmt.Println("主板ID:", board)
	cpu, _ := GetCpuId()
	fmt.Println("CPUID:", cpu)
	UUID, _ := GetSysUUID()
	fmt.Println("UUID:", UUID)
	bios, _ := GetBIOSID()
	fmt.Println("BIOS序列号:", bios)
}
