package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	for x, c := range base {
		for y, c1 := range base {
			if len([]rune(base)) <= 1 || (c == c1 && x != y) || !IsAlpha(string(c1)) {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	negative := false
	if nbr < 0 {
		negative = true
		nbr = nbr * -1
	}
	nbase := len([]rune(base))
	str := ""
	bound := 0
	for nbr/nbase > 0 {
		bound++
		if bound == 5000 {
			return
		}
		str = string(base[nbr%nbase]) + str
		nbr = nbr / nbase
	}
	if negative {
		z01.PrintRune('-')
	}
	z01.PrintRune(rune(base[nbr]))
	for _, s := range str {
		z01.PrintRune(s)
	}
}
