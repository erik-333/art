package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func multiLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	fmt.Println("Enter your text, press Enter twice to finish")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// Only remove escape sequences while preserving other characters
		// Preserve the square bracket
		line = strings.ReplaceAll(line, "\x1b", "[")
		// Remove any remaining escape characters
		line = strings.ReplaceAll(line, "\x1b", "")

		// Dont trim spaces as they might be part of the input
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	result := strings.Join(lines, "\n")
	return result
}
