package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return // if x or y is less than 0 do nothing
	}
	// Looping through the height
	for height := 0; height < y; height++ {
		// Looping through the width
		for width := 0; width < x; width++ {
			// Mapping out the corners
			if height == 0 && width == 0 {
				z01.PrintRune('A')
			} else if height == y-1 && width == 0 {
				z01.PrintRune('C')
			} else if height == 0 && width == x-1 {
				z01.PrintRune('C')
			} else if height == y-1 && width == x-1 {
				z01.PrintRune('A')
			} else if height == 0 || height == y-1 {
				z01.PrintRune('B') // Mapping out the height
			} else if width == 0 || width == x-1 {
				z01.PrintRune('B') // Mapping out the width
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')

	}
}

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program_name width height")
		return
	}

	// Parse command-line arguments using strconv
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

	// Call the QuadD function with command-line argument values
	QuadE(width, height)
}
