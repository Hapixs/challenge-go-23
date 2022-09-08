package piscine

import "github.com/01-edu/z01"

func StrRev(s string) {
	runes := []rune{}

	for _, st := range s {
		runes = append(runes, st)
	}
	for i := len(runes); i >= 0; i-- {
		z01.PrintRune(runes[i])
	}
}
