package main

import (
	"os"
)

func main() {
	found := false
	for _, arg := range os.Args[1:] {
		if arg == "01" || arg == "galaxy" || arg == "galaxy01" {
			found = true
			break
		}
	}
	if found {
		// Do nothing if found
	}
}
