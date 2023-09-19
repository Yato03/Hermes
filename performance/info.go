package performance

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"Hermes/utils"
)

func CpuInfo() string {

	var result = ""

	// Obtener información de CPU
	cpuInfo, _ := cpu.Info()

	fmt.Println("CPU Info:")
	for _, info := range cpuInfo {
		result += fmt.Sprintf("  %s: %d\n", info.ModelName, info.Cores)
	}

	result += "\n"

	cpuUsage, err := utils.CpuUsage()
	cpuUsageString := fmt.Sprintf("%.2f", cpuUsage)

	if err != nil {
		result += fmt.Sprintf("It couldn't be possible to obtain usage information")
	} else {
		result += fmt.Sprintf("CPU Usage: %s", cpuUsageString)
	}

	return result
}

func MemInfo() string {
	var result = ""

	memInfo, _ := mem.VirtualMemory()

	memTotal := memInfo.Total / 1024 / 1024 / 1024

	memUsed := memInfo.Used / 1024 / 1024 / 1024
	memUsedPercent := memInfo.UsedPercent

	memFree := memInfo.Free / 1024 / 1024 / 1024
	memFreePercent := float64(memFree) / float64(memTotal) * 100

	memAvaiable := memInfo.Free / 1024 / 1024 / 1024
	memAvaiablePercent := float64(memAvaiable) / float64(memTotal) * 100

	result += fmt.Sprintf("Used memory: %vGB/%vGB (%.2f%%) \n", memUsed, memTotal, memUsedPercent)
	result += fmt.Sprintf("Free memory: %vGB/%vGB (%.2f%%) \n", memFree, memTotal, memFreePercent)
	result += fmt.Sprintf("Free memory: %vGB/%vGB (%.2f%%) \n", memAvaiable, memTotal, memAvaiablePercent)

	return result
}

func DiskInfo() (string, error) {
	var result = ""

	partitions, err := utils.DiskUsage()
	if err != nil {
		return "", err
	}

	for _, partition := range partitions {
		result += "\n"
		result += fmt.Sprintf("Device: %s\n", partition.Device)
		result += fmt.Sprintf("Total space: %.2f GB\n", partition.Total)
		result += fmt.Sprintf("\tUsed space: %.2f GB\n", partition.Used)
		result += fmt.Sprintf("\tAvaiable space: %.2f GB\n", partition.Available)
	}
	return result, nil
}

func NetInfo() string {
	var result = ""

	return result
}

func SysInfo() string {
	var result = ""

	return result
}
