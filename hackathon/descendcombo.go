package tests

import "fmt"

func DescendComb() {
	for i := 99; i >= 10; i-- {
		for j := i - 1; j >= 10; j-- {
			fmt.Print(i)
			fmt.Print(" ")
			fmt.Print(j)
			if i != 10 || j != 10 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}
