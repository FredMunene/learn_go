package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[2:]
	for _, str := range arg {
		for _, cha := range str {
			z01.PrintRune(cha)
		}
		z01.PrintRune('\n')
	}
}
