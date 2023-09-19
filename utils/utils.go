package utils

import (
	"math"
	"strconv"
)

func ParseUint64(s string) uint64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return uint64(n)
	} else {
		return 0
	}

}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func ByToGb(n uint64) float64 {
	return float64(n) * math.Pow(10, -9)
}
