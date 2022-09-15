package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	for _, c := range base {
		for _, c1 := range base {
			if c == c1 {
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
	for nbr/nbase > 0 {
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
