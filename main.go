package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ptrn = regexp.MustCompile(`\[(\d+) ([^\]]+)]`)

// decode function expands patterns like [5 #] into #####
func decode(input string) (string, error) {
	if strings.Count(input, "[") != strings.Count(input, "]") {
		return "", fmt.Errorf("Error")
	}

	result := ptrn.ReplaceAllStringFunc(input, func(match string) string {
		matches := ptrn.FindStringSubmatch(match)
		if len(matches) != 3 {
			return "Error"
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil || count <= 0 {
			return "Error"
		}

		for _, isdigit := range matches[1] {
			if isdigit < '0' || isdigit > '9' {
				return "Error"
			}

		}

		return strings.Repeat(matches[2], count)
	})

	if strings.Contains(result, "Error") {
		return "", fmt.Errorf("Error")
	}

	return result, nil
}

// encode function compresses repeating sequences into [n seq] format
func encode(input string) string {
	var result strings.Builder
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		j := i
		for j < n && runes[j] == runes[i] {
			j++
		}

		count := j - i
		if count > 3 {
			result.WriteString(fmt.Sprintf("[%d %s]", count, string(runes[i])))
		} else {
			result.WriteString(strings.Repeat(string(runes[i]), count))
		}
		i = j
	}

	return result.String()
}

func main() {
	doubleFlag := flag.Bool("2x", false, "double flag")
	decodeFlag := flag.Bool("de", false, "decode mode")
	encodeFlag := flag.Bool("en", false, "encode mode")
	helpFlag := flag.Bool("h", false, "help")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error")
		fmt.Println("Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if u want to double it use -2x after -de|-en")
		return
	}

	input := args[0]

	if *helpFlag {
		fmt.Println("Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if u want to double it use -2x after -de|-en")
		os.Exit(0)
	}

	if *doubleFlag {
		// Double digits in pattern [n X]
		input = ptrn.ReplaceAllStringFunc(input, func(match string) string {
			matches := ptrn.FindStringSubmatch(match)
			if len(matches) != 3 {
				return match
			}
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				return match
			}
			return fmt.Sprintf("[%d %s]", num*2, matches[2])
		})
		
		// Double characters outside of patterns
		segments := ptrn.Split(input, -1)
		for i, seg := range segments {
			segments[i] = strings.Repeat(seg, 2)
		}
	
		// Reconstruct the string with patterns
		matches := ptrn.FindAllStringSubmatch(input, -1)
		result := segments[0]
		for i := 0; i < len(matches); i++ {
			if i < len(segments)-1 {
				result += fmt.Sprintf("[%s %s]", matches[i][1], matches[i][2]) + segments[i+1]
			}
		}
		input = result
	}
	

	if *decodeFlag {
		result, err := decode(input)
		if err != nil {
			fmt.Println("Error")
			return
		}
		fmt.Println(result)
	} else if *encodeFlag {
		fmt.Println(encode(input))
	} else {
		fmt.Println("Error")
		fmt.Println("Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if u want to double it use -2x after -de|-en")
		os.Exit(1)
	}
}