package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	red    = "\033[31m"
	yellow = "\033[33m"
	purple = "\033[35m"
	reset  = "\033[0m"
)

func main() {
	doubleFlag := flag.Bool("2x", false, "double flag")
	decodeFlag := flag.Bool("de", false, "decode mode")
	encodeFlag := flag.Bool("en", false, "encode mode")
	multiflag := flag.Bool("m", false, "multiline")
	helpFlag := flag.Bool("h", false, "help")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 && !*multiflag && !*helpFlag {
		fmt.Println(red + "Error! :O\n" + reset)
		fmt.Println(yellow + "Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if want to double size, use -2x")
		fmt.Println("for multiline input, use -m" + reset)
		return
	}

	if *helpFlag {
		fmt.Println(yellow + "Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if want to double size, use -2x")
		fmt.Println("for multiline input, use -m" + reset)
		os.Exit(0)
	}

	var input string
	if *multiflag {
		input = multiLine()
		// check if any input
		if len(input) == 0 || (len(input) == 1 && input == "") {
			fmt.Println(red + "No input provided\n" + reset)
			return
		}

	} else if len(args) > 0 {
		input = args[0]
	}

	if *doubleFlag {
		doubled := doubleASCIIArt([]string{input})
		result := doubled
		if *encodeFlag {
			// Encode each line if encode flag is set
			result = make([]string, len(doubled))
			for i, line := range doubled {
				result[i] = encode(line)
			}
		}
		fmt.Printf("\n")
		fmt.Println(strings.Join(result, "\n"))
	}

	if *decodeFlag && !*doubleFlag {
		if *multiflag {
			lines := strings.Split(input, "\n")
			var decodedLines []string

			// check if any input
			if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
				fmt.Println(red + "No input provided\n" + reset)
				return
			}

			for i, line := range lines {
				// skip empty lines
				if strings.TrimSpace(line) == "" {
					continue
				}

				decoded, err := decode(line)
				if err != nil {
					// debug info
					fmt.Printf(red+"Failed to decode line %d: %q\n"+reset, i+1, line)
					fmt.Printf(red+"Error! :O , details: %v\n"+reset, err)
					return
				}
				decodedLines = append(decodedLines, decoded)
			}
			if len(decodedLines) > 0 {
				fmt.Printf("\n")
				fmt.Println(strings.Join(decodedLines, "\n"))
			} else {
				fmt.Println(red + "No valid lines found\n" + reset)
				return
			}
		} else {
			if strings.TrimSpace(input) == "" {
				fmt.Println(red + "Error! :0 , no input\n" + reset)
				return
			}
			result, err := decode(input)
			if err != nil {
				fmt.Printf(red+"Failed to decode: %q\n\n"+reset, input)
				fmt.Printf(red+"Error! :O , details: %v\n"+reset, err)
				return
			}
			fmt.Printf("\n")
			fmt.Println(result)
		}
	} else if *encodeFlag && !*doubleFlag {
		fmt.Printf("\n")
		fmt.Println(encode(input))
	} else if !*decodeFlag && !*encodeFlag && !*doubleFlag && !*multiflag {
		fmt.Println(red + "Error! :O\n" + reset)
		fmt.Println(yellow + "Usage: go run . -de|-en <text>")
		fmt.Println("-de flag is for decode mode and -en is for encode mode, if want to double size, use -2x")
		fmt.Println("for multiline input, use -m" + reset)
		os.Exit(1)
	}
}
