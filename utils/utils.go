package utils

import (
	"math"
	"strconv"
	"unicode"
)

func ParseUint64(s string) uint64 {

	//first character is not a digit or sign
	firstChar := rune(s[0])
	if !unicode.IsDigit(firstChar) && !(firstChar == '+' || firstChar == '-') {
		return 0
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	} else {
		return uint64(n)
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
