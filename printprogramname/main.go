package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0][2:]
	for _, char := range args {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
