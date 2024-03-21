package tests

import (
	"fmt"
)

func EightQueens() {

}

const N = 9

func Grid(int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Print(i)
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func IsSafe(row, col, num int) {
	//for i := 2 {}

}
