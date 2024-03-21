package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for height := 0; height < y; height++ {

		for width := 0; width < x; width++ {

			if height == 0 && width == 0 {
				z01.PrintRune('/')
			} else if height == y-1 && width == 0 {
				z01.PrintRune('\\')
			} else if height == 0 && width == x-1 {
				z01.PrintRune('\\')
			} else if height == y-1 && width == x-1 {
				z01.PrintRune('/')
			} else if height == 0 || height == y-1 {
				z01.PrintRune('*')
			} else if width == 0 || width == x-1 {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')

	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program_name width height")
		return
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: Invalid width argument '%s'. Please provide an integer.\n", os.Args[1])
		return
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: Invalid height argument '%s'. Please provide an integer.\n", os.Args[2])
		return
	}

	QuadB(width, height)
}
