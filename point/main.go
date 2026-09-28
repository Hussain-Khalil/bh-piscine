package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printerone(s string) {
	arr := []rune(s)
	for _, r := range arr {
		z01.PrintRune(r)
	}
}

func printertwo(r rune) {
	z01.PrintRune(r)
}

func main() {
	points := point{}
	setPoint(&points)
	masx := "x = "
	masy := ", y = "
	printerone(masx)
	printertwo(52)
	printertwo(50)
	printerone(masy)
	printertwo(50)
	printertwo(49)
	printertwo(10)
}
