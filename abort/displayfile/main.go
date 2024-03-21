package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("File name missing")
	} else if len(args) == 1 {
		fmt.Println(args[0])
	} else {
		fmt.Println("Too many arguments")
	}
}
