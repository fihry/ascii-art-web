package main

import (
	"os"
	"strings"
	"testing"

	"ascii"
)

func TestRunWithArgs(t *testing.T) {
	// Define input-output map
	tests := map[string]func() map[string]string{
		"standard":   ascii.Standard,
		"shadow":     ascii.Shadow,
		"thinkertoy": ascii.Thinkertoy,
	}
	// Iterate over the tests
	for banner, test := range tests {
		// Run the test
		test := test()
		for text, expectedOutput := range test {
			// Run the program and capture the output
			output := ascii.AsciiArt(text, banner)
			// Compare the output with the expected output
			if strings.TrimSpace(output) != strings.TrimSpace(ascii.RemoveTrailingSpaces(expectedOutput)) {
				os.Stdout.WriteString("\033[33mtests :" + text + " -> Failed!\n\033[0m")
				t.Errorf("\033[31mOutput mismatch.\033[0m \n========================================= \033[31mEXPECTED\033[0m =========================================\n%s\n=========================================   \033[31mGOT\033[0m   =========================================\n%s", expectedOutput, output)
			} else {
				os.Stdout.WriteString("\033[32mtests :" + text + " -> passed!\n\033[0m")
			}
		}
	}
}
