package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	isUpper := args[0] == "--upper"
	for _, c := range args {
		val := Atoi(c)
		if val == -1 || val > 26 {
			z01.PrintRune(' ')
		} else if val >= 0 {
			z01.PrintRune(map[bool]rune{true: rune(val + 64), false: rune(val + 96)}[isUpper])
		}
	}
	z01.PrintRune('\n')
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	total := 0
	if s[0] == '-' {
		return -2
	}
	if !IsNumeric(s) {
		return -1
	}
	for _, c := range s {
		total = total*10 + int(c-48)
	}
	return total
}
