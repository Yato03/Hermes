package performance

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func CpuInfo() string {

	var result = ""

	// Obtener información de CPU
	cpuInfo, _ := cpu.Info()

	fmt.Println("CPU Info:")
	for _, info := range cpuInfo {
		result += fmt.Sprintf("  %s: %d\n", info.ModelName, info.Cores)
	}
	return result
}

func MemInfo() string {
	var result = ""

	memInfo, _ := mem.VirtualMemory()

	memTotal := memInfo.Total / 1024 / 1024 / 1024

	memUsed := memInfo.Used / 1024 / 1024 / 1024
	memUsedPercent := memInfo.UsedPercent

	memFree := memInfo.Available / 1024 / 1024 / 1024
	memFreePercent := float64(memFree) / float64(memTotal) * 100

	result += fmt.Sprintf("Used memory: %vGB/%vGB (%.2f%%) \n", memUsed, memTotal, memUsedPercent)
	result += fmt.Sprintf("Free memory: %vGB/%vGB (%.2f%%) \n", memFree, memTotal, memFreePercent)

	return result
}

func DiskInfo() string {
	var result = ""

	return result
}

func NetInfo() string {
	var result = ""

	return result
}

func SysInfo() string {
	var result = ""

	return result
}
