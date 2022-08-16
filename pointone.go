package piscine

import (
	"github.com/01-edu/z01"
)

func PointOne(n *int) {
	n := 0

	p := &n

	*p = 1

	z01.PrintRune(n)
}
