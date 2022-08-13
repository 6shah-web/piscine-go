package piscine

import "github.com/01-edu/z01"

func a(r rune) {
	z01.PrintRune(r)
}

func IsNegative(nb int) {
	if nb < 0 {
		a(84)
	} else {
		a(70)
	}
	a('\n')
}
