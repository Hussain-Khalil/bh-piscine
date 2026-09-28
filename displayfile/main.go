package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else if len(args) < 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Almost there!!")
	}
}
