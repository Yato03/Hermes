package utils

import (
	"errors"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

type DiskInfo struct {
	Device    string
	Total     string
	Used      string
	Available string
}

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

func DiskUsage() ([]DiskInfo, error) {
	cmd := exec.Command("df", "-h")
	output, err := cmd.CombinedOutput()
	if err != nil {
		errorMessage := fmt.Sprintf("Error al ejecutar df:%s", err)
		return []DiskInfo{}, errors.New(errorMessage)
	}

	// Dividir la salida en líneas
	lines := strings.Split(string(output), "\n")

	devices := []DiskInfo{}

	// Ignorar la primera línea (encabezado)
	for i, line := range lines {
		if i == 0 {
			continue
		}

		// Dividir la línea en campos
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		// Obtener los campos de interés
		device := fields[0]
		total := fields[1]
		used := fields[2]
		available := fields[3]

		devices = append(devices, DiskInfo{device, total, used, available})
	}
	return devices, nil
}

func parseUint64(s string) uint64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return uint64(n)
	} else {
		return 0
	}

}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func ByToGb(n uint64) float64 {
	return float64(n) * math.Pow(10, -9)
}
