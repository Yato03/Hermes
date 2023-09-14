package utils

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func CpuUsage() (float64, error) {
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
			cpuUsageFloat := parseFloat(cpuUsage)
			totalCPU += cpuUsageFloat
		}
	}

	return totalCPU, nil

}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}
