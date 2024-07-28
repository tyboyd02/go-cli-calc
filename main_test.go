package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runCLI(args ...string) (string, error) {
	cmd := exec.Command("./go-cli-calc", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestCalculator(t *testing.T) {
	tests := []struct {
		number     string
		expression string
		expected   string
	}{
		{"1", "2 + 2", "4"},
		{"2", "3 * 3", "9"},
		{"3", "6 / 2", "3"},
		{"4", "5 - 3", "2"},
		{"5", "2 ^ 3", "8"},
		{"6", "(2 + 3) * 4", "20"},
		{"7", "2 + (3 * 4)", "14"},
		{"8", "(2 + 3) / (1 + 1)", "2.5"},
		{"9", "2 + 3 * 4 ^ 2", "50"},
		{"10", "(2 + 3) * (4 - 1)", "15"},
		{"11", "2 * (3 + 4) / (1 + 1)", "7"},
		{"12", "(2 ^ 3) * (4 - 1)", "24"},
		{"13", "(2 + 3) * (4 / (1 + 1))", "10"},
		{"14", "2 ^ (3 * 2)", "64"},
		{"15", "(2 ^ 3) ^ 2", "64"},
		{"16", "2 + 3 * 4 / 2", "8"},
		{"17", "(2 + (3 * 4)) / 2", "7"},
		{"18", "((2 + 3) * 4) / 2", "10"},
		{"19", "2 + (3 + 4) * (5 - 1)", "30"},
		{"20", "10 / (2 + 3)", "2"},
		{"21", "(10 / 2) + 3", "8"},
		{"22", "(10 - 2) * 3", "24"},
		{"23", "2 * (3 + 4) - 5", "9"},
		{"24", "10 - (2 ^ 2)", "6"},
		{"25", "10^2 - -6 +-3 * (9*2^2-3)", "7"},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			output, err := runCLI(test.expression)
			if err != nil {
				t.Fatalf("Error running CLI command: %v", err)
			}
			output = strings.TrimSpace(output)
			if output != test.expected {
				t.Errorf("For test #%s, expected %q but got %q", test.number, test.expected, output)
			}
		})
	}
}
