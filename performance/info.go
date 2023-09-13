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

func MemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("\nTotal de memoria: %v GB\n", memInfo.Total/1024/1024/1024)
	fmt.Printf("Memoria libre: %v GB\n", memInfo.Free/1024/1024/1024)
}
