package main

import (
	"fmt"
	"os"
)

func main() {
	for _, argument := range os.Args {
		fmt.Println(argument)
	}
}
