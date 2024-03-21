package main

import (
	"fmt"
	tests "tests/hackathon"
)

func main() {
	a := []int{1, 1, 2, 3, 4, 3, 4}
	unmatch := tests.Unmatch(a)
	fmt.Println(unmatch)
}
