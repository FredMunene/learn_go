package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Check if the --upper flag is provided
	upperCase := false
	if len(args) > 0 && args[0] == "--upper" {
		upperCase = true
		args = args[1:] // Remove the flag from arguments
	}

	for _, arg := range args {
		n := atoi(arg) // Manual conversion to integer
		if n == -1 {
			z01.PrintRune(' ') // Print space for non-integer arguments
			continue
		}

		// Check if n is within the valid range of alphabet positions
		if n >= 1 && n <= 26 {
			var letter rune
			if upperCase {
				letter = 'A' + rune(n-1)
			} else {
				letter = 'a' + rune(n-1)
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ') // Print space for out-of-range positions
		}
	}

	if len(args) > 0 {
		z01.PrintRune('\n') // Print newline after processing all arguments if there are arguments
	}
}

// atoi manually converts a string to an integer
func atoi(s string) int {
	var result int
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1 // Return -1 for non-integer strings
		}
		result = result*10 + int(ch-'0')
	}
	return result
}
