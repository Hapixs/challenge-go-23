package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	for x, c := range base {
		for y, c1 := range base {
			if len([]rune(base)) <= 1 || (c == c1 && x != y) || c1 == '-' || c1 == '+' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
	}
	recPrintBase(nbr, len([]rune(base)), base)
}

func recPrintBase(nbr int, nbase int, base string) {
	if nbr/nbase == 0 {
		z01.PrintRune(rune(base[printBaseAbs(nbr)]))
	} else {
		recPrintBase(nbr/nbase, nbase, base)
		z01.PrintRune(rune(base[printBaseAbs(nbr%nbase)]))
	}
}

func printBaseAbs(f int) int {
	return map[bool]int{true: f * -1, false: f}[f < 0]
}
