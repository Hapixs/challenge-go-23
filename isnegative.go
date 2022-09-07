package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	z01.PrintRune(map[bool]rune{true: 'T', false: 'F'}[nb < 0])
	z01.PrintRune('\n')
}
