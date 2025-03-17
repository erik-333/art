package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// pattern for the actual replacement
var ptrn = regexp.MustCompile(`\[(\d+) ([^\]]+)]`)

// pattern for validation and error handling
// old pattern below
// var validatePtrn = regexp.MustCompile(`\[(\S+)\s+(.+?)\]`)

// New and improved pattern
var validatePtrn = regexp.MustCompile(`\[(\S+)(?:\s*)?\]`)

// pattern for empty 2nd arg checking (eg. [2 ])
var emptySecondArgPattern = regexp.MustCompile(`\[\d+\s+\]`)

// pattern for input without space
var noSpacePattern = regexp.MustCompile(`\[\d+[^\s\]]+\]`)

// pattern for single argument (eg. [2])
var singleArgPattern = regexp.MustCompile(`\[\d+\]`)

func decode(input string) (string, error) {
	// remove all escape characters from the input
	input = strings.TrimPrefix(input, "\x1b")

	// check for mismatched brackets
	if strings.Count(input, "[") != strings.Count(input, "]") {
		return "", fmt.Errorf("%s", red+"Error! :O , mismatched brackets in input"+reset)
	}

	// Check if the input contains a pattern without space between arguments
	if noSpacePattern.MatchString(input) {
		return "", fmt.Errorf("%s", red+"Error! :O , arguments not separated by space"+reset)
	}

	// Check for empty second argument
	if emptySecondArgPattern.MatchString(input) {
		return "", fmt.Errorf("%s", red+"Error! :O , empty second argument"+reset)
	}

	if singleArgPattern.MatchString(input) {
		return "", fmt.Errorf("%s", red+"Error! :O , empty second argument"+reset)
	}

	// first do validation using validatePtrn
	matches := validatePtrn.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if _, err := strconv.Atoi(match[1]); err != nil {
			return "", fmt.Errorf("%s", red+"Error! :O , first argument must be a number\n"+reset)
		}
	}

	// If validation passes, proceed with the actual replacement using ptrn
	result := ptrn.ReplaceAllStringFunc(input, func(match string) string {
		matches := ptrn.FindStringSubmatch(match)
		if len(matches) != 3 {
			fmt.Println(red + "Error! :O , invalid pattern format" + reset)
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Println(red + "Error! :O , invalid number format" + reset)
		}

		if count <= 0 {
			fmt.Println(red + "Error! :O , count must be positive" + reset)
		}

		return strings.Repeat(matches[2], count)
	})

	if strings.Contains(result, "Error:") {
		return "", errors.New(result)
	}

	return result, nil
}
