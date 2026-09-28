package piscine

import "github.com/01-edu/z01"

func printer(c rune) {
	z01.PrintRune(c)
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printer(rune(i/10 + '0'))
			printer(rune(i%10 + '0'))
			printer(' ')
			printer(rune(j/10 + '0'))
			printer(rune(j%10 + '0'))
			if i == 1 && j == 0 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
