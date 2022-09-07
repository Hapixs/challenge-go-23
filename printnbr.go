package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	s1 := strconv.Itoa(n)
	for _, ch := range s1 {
		z01.PrintRune(ch)
	}
}
