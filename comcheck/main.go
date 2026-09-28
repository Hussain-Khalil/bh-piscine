package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	found := false
	for _, chr := range args {
		if chr == "01" || chr == "galaxy" || chr == "galaxy 01" {
			found = true
		}
	}
	if found {
		fmt.Println("Alert!!!")
	}
}
