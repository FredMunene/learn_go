package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else if len(args) == 1 {
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for len(content) > 0 && (content[len(content)-1] == '\n' || content[len(content)-1] == '\r') {
			content = content[:len(content)-1]
		}
		fmt.Println(string(content))
	}
}
