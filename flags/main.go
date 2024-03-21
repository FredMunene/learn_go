package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var insertString string
	var baseString string
	var orderFlag bool

	// Parse arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--insert", "-i":
			if i+1 < len(args) {
				insertString = args[i+1]
				i++ // Move to the next argument
			}
		case "--order", "-o":
			orderFlag = true
		default:
			baseString = arg
		}
	}

	// Apply modifications
	result := baseString // Get the base string
	if insertString != "" {
		// Insert the string after the first character occurrence
		firstCharIndex := strings.IndexFunc(result, func(r rune) bool {
			return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
		})

		if firstCharIndex != -1 {
			result = result[:firstCharIndex+1] + insertString + result[firstCharIndex+1:]
		} else {
			result += insertString // Append if no character is found
		}
	}

	// Order the result if the --order flag is provided
	if orderFlag {
		result = orderString(result)
	}

	// Print the modified string
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func orderString(s string) string {
	slice := []rune(s)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return string(slice)
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t\t This flag will behave like a boolean, if it is called it will order the argument.")
}
