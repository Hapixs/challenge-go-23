package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
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
