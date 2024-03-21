package piscine

import (
	"github.com/01-edu/z01"
)

// func prints in descending order and on a single line :all posible combinations of two different two digit numbers
func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '8'; j >= i; j-- { // j will always be different from j
			z01.PrintRune(i)
			z01.PrintRune(i)
			z01.PrintRune(' ')
			z01.PrintRune(j)
			z01.PrintRune(j)
			if i == '8' && j == '9' {
				z01.PrintRune(',')
			}
		}
	}
}
