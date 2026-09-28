package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args) - 1; i > 0; i-- {
		for _, cha := range args[i] {
			z01.PrintRune(cha)
		}
		if i != 0 {
			z01.PrintRune('\n')
		}
	}
}
