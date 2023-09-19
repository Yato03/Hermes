package performance

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type DeviceInfo struct {
	Device    string
	Total     string
	Used      string
	Available string
}

func DiskInfo() (string, error) {
	var result = ""

	partitions, err := diskUsage()
	if err != nil {
		return "", err
	}

	for _, partition := range partitions {
		result += "\n"
		result += fmt.Sprintf("Device: %s\n", partition.Device)
		result += fmt.Sprintf("Total space: %s\n", partition.Total)
		result += fmt.Sprintf("\tUsed space: %s\n", partition.Used)
		result += fmt.Sprintf("\tAvaiable space: %s\n", partition.Available)
	}
	return result, nil
}

func diskUsage() ([]DeviceInfo, error) {
	cmd := exec.Command("df", "-h")
	output, err := cmd.CombinedOutput()
	if err != nil {
		errorMessage := fmt.Sprintf("Error al ejecutar df:%s", err)
		return []DeviceInfo{}, errors.New(errorMessage)
	}

	// Dividir la salida en líneas
	lines := strings.Split(string(output), "\n")

	devices := []DeviceInfo{}

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

		devices = append(devices, DeviceInfo{device, total, used, available})
	}
	return devices, nil
}
