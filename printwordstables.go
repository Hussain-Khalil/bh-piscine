package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, world := range a {
		for _, charact := range world {
			z01.PrintRune(charact)
		}
		z01.PrintRune('\n')
	}
}
