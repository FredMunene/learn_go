package main

import (
	"fmt"
	"tests"

	"github.com/01-edu/z01"
)

func main() {
	result := tests.Rot14("a6 eY  ke  2I")
	fmt.Print(result)
	z01.PrintRune('\n')

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
