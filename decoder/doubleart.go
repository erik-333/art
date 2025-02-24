package main

import (
	"regexp"
	"strconv"
)

// doubleNumber doubles values found in the input string
func doubleNumber(input string) string {
	re := regexp.MustCompile(`\d+`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		num, _ := strconv.Atoi(match)
		return strconv.Itoa(num * 2)
	})
}

func doubleASCIIArt(input []string) []string {
	// First decode all lines
	decodedLines := make([]string, 0, len(input))
	for _, line := range input {
		decoded, err := decode(line)
		if err != nil {
			continue
		}
		decodedLines = append(decodedLines, decoded)
	}

	// Then double the decoded lines
	result := make([]string, 0, len(decodedLines))
	for _, line := range decodedLines {
		var doubled string
		for _, char := range line {
			doubled += string(char) + string(char)
		}
		result = append(result, doubled)
	}

	return result
}
