package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := len(arg) - 1; i >= 0; i-- {
		for _, cha := range arg[i] {
			z01.PrintRune(cha)
		}
		z01.PrintRune('\n')
	}
}
