package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: No input provided\n")
		return
	}
	if len(os.Args) > 2 {
		fmt.Printf("Error: Too many arguments provided\n")
		return
	}
	input := os.Args[1]
	fmt.Println("Input:", input)
	artdecode(input)
	fmt.Printf("^^ Decoded input ^^ \n")
}

func artdecode(input string) string {
	ptrn := regexp.MustCompile("\\[(\\d+) (.+?)\\]")
	match := ptrn.FindStringSubmatch(input)

	if match == nil {
		fmt.Printf("Error \n")
		return ""
	}

	count, err := strconv.Atoi(match[1])
	if err != nil || count == 0 {
		fmt.Printf("Error \n")
		return ""
	}

	result := strings.Repeat(match[2], count)

	fmt.Println(result)
	return ""
}
