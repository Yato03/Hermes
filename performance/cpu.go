package performance

import (
	"Hermes/utils"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
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

	cpuUsage, err := cpuUsage()
	cpuUsageString := fmt.Sprintf("%.2f", cpuUsage)

	if err != nil {
		result += fmt.Sprintf("It couldn't be possible to obtain usage information")
	} else {
		result += fmt.Sprintf("CPU Usage: %s", cpuUsageString)
	}

	return result
}

func cpuUsage() (float64, error) {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return 0, errors.New("Error executing ps aux")
	}

	lines := strings.Split(string(output), "\n")

	totalCPU := 0.0

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			cpuUsage := fields[2]
			cpuUsageFloat := utils.ParseFloat(cpuUsage)
			totalCPU += cpuUsageFloat
		}
	}

	return totalCPU, nil

}
