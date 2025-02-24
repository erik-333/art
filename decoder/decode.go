package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"errors"
)

// old semi-working pattern regex
var ptrn = regexp.MustCompile(`\[(\d+) ([^\]]+)]`)

//updated regex, revert to old if any issues occur
// var ptrn = regexp.MustCompile(`\[(\S+)\s+(.+?)\]`)


func decode(input string) (string, error) {
    // Remove all escape characters from the input
    input = strings.TrimPrefix(input, "\x1b")
    
    // Check for matching brackets
    if strings.Count(input, "[") != strings.Count(input, "]") {
        return "", fmt.Errorf("mismatched brackets in input")
    }

    result := ptrn.ReplaceAllStringFunc(input, func(match string) string {
        matches := ptrn.FindStringSubmatch(match)
        if len(matches) != 3 {
            return "Error: invalid pattern format"
        }

		if !IsNumeric(matches[1]) {
			fmt.Println("Error: first argument must be a number")
		}

        count, err := strconv.Atoi(matches[1])
        if err != nil {
            fmt.Println("Error: invalid number format")
        }
        
        if count <= 0 {
            fmt.Println("Error: count must be positive")
        }

        // No need for additional digit checking since strconv.Atoi already validates this
        return strings.Repeat(matches[2], count)
    })

    if strings.Contains(result, "Error:") {
		return "", errors.New(result)
    }

    return result, nil
}

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, n := range s {
		if n >= '0' && n <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}