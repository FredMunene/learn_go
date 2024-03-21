package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg)-1; i++ {
		for j := 0; j < len(arg)-i-1; j++ {
			if arg[j] > arg[j+1] {
				// Swap elements
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	for _, str := range arg {
		for _, cha := range str {
			z01.PrintRune(cha)
		}
		z01.PrintRune('\n')
	}
}
