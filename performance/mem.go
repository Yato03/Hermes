package performance

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

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
