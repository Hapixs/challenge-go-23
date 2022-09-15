package piscine

import "github.com/01-edu/z01"

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
	negative := false
	if nbr < 0 {
		negative = true
	}
	nbase := len([]rune(base))
	str := ""
	for nbr/nbase != 0 {
		index := nbr % nbase
		if index < 0 {
			index = index * -1
		}
		str = string(base[index]) + str
		nbr = nbr / nbase
	}
	if negative {
		z01.PrintRune('-')
	}
	if nbr < 0 {
		nbr = nbr * -1
	}
	z01.PrintRune(rune(base[nbr]))
	for _, s := range str {
		z01.PrintRune(s)
	}
}
