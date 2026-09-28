package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := range args {
		if i > 0 {
			for _, cha := range args[i] {
				z01.PrintRune(cha)
			}
		}
		if i > 0 {
			z01.PrintRune('\n')
		}
	}
}
