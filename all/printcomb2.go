package piscine

import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ { // j will always be different from j
			fmt.Printf("%02d %02d", i, j) // printed with leading zeros using %02d
			if !(i == 98 && j == 99) {
				fmt.Printf(", ")
			}
		}
	}
	fmt.Println()
}
