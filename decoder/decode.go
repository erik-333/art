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
var validatePtrn = regexp.MustCompile(`\[(\S+)\s+(.+?)\]`)

func decode(input string) (string, error) {
	// remove all escape characters from the input
	input = strings.TrimPrefix(input, "\x1b")

	// check for mismatched brackets
	if strings.Count(input, "[") != strings.Count(input, "]") {
		return "", fmt.Errorf(red + "Error! :O , mismatched brackets in input" + reset)
	}

	// first do validation using validatePtrn
	for _, match := range validatePtrn.FindAllStringSubmatch(input, -1) {
		if len(match) == 3 {
			if _, err := strconv.Atoi(match[1]); err != nil {
				return "", fmt.Errorf(red + "Error! :O , first argument must be a number\n" + reset)
			}
		}
	}

	// If validation passes, proceed with the actual replacement using ptrn
	result := ptrn.ReplaceAllStringFunc(input, func(match string) string {
		matches := ptrn.FindStringSubmatch(match)
		if len(matches) != 3 {
			return red + "Error! :O , invalid pattern format" + reset
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
