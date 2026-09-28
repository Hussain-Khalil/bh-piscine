package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	str := []rune(args[0])
	for _, char := range str {
		if char != '.' && char != '/' {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
