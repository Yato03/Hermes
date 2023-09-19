package utils

import (
	"fmt"
	"testing"
)

func TestParseUint64(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"12345", 12345},
		{"0", 0},
		{"9223372036854775807", 9223372036854775807}, // Max value int64
		{"-12345", 0},
		{"abc", 0},
		{"", 0},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ParseUint64(test.input)
			if result != test.expected {
				t.Errorf("Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"3.14159", 3.14159},
		{"0", 0.0},
		{"-2.71828", -2.71828},
		{"abc", 0.0},
		{"", 0.0},
		{"1.23e-5", 1.23e-5},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ParseFloat(test.input)
			if result != test.expected {
				t.Errorf("Input: %s, Expected: %f, Got: %f", test.input, test.expected, result)
			}
		})
	}
}

func TestByToGb(t *testing.T) {
	tests := []struct {
		input    uint64
		expected float64
	}{
		{1024, 1e-9},            // 1 Kilobyte to Gigabytes
		{1073741824, 1.0},       // 1 Gigabyte to Gigabytes
		{5368709120, 5.0},       // 5 Gigabytes to Gigabytes
		{1000000000000, 1000.0}, // 1 Terabyte to Gigabytes
		{1234567890123456789, 1234.5678901234568},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d to GB", test.input)
		t.Run(testName, func(t *testing.T) {
			result := ByToGb(test.input)
			if result != test.expected {
				t.Errorf("Input: %d, Expected: %f, Got: %f", test.input, test.expected, result)
			}
		})
	}
}
