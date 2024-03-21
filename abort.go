package piscine

import "fmt"

func Abort(a, b, c, d, e int) int {
	arry := []int{a, b, c, d, e}
	// bubble sort algorithm
	for i := 0; i < len(arry); i++ { // iterates through slice
		for j := i + 1; j < len(arry)-i-1; j++ { // runs over unsorted part
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
	}
	fmt.Println(arry)
	MedIndex := (len(arry) / 2)
	fmt.Println(MedIndex)
	return (arry[MedIndex])
}
