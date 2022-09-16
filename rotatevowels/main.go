package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := Join(os.Args[1:], " ")
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		a := r[i]
		if isVowel(a) {
			for y := i; y < len(r); y++ {
				b := r[y]
				if isVowel(b) {
					r[i], r[y] = r[y], r[i]
				}
			}
		}
	}
	for _, c := range r {
		z01.PrintRune(c)
	}
	if len(r) > 0 {
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func Join(s []string, sep string) string {
	if len(s) <= 0 {
		return ""
	}
	return s[0] + map[bool]string{true: sep, false: ""}[len(s) > 1] + Join(s[1:], sep)
}

func isVowel(r rune) bool {
	str := "aeiouAEIOU"
	return StrIndex(str, string(r)) >= 0
}

func StrIndex(s, find string) int {
	SizeS := StrLen(s)
	SizeF := StrLen(find)
	for i := 0; i < SizeS-SizeF; i++ {
		if s[i:i+SizeF] == find {
			return i
		}
	}
	return -1
}

func StrLen(s string) int {
	return len([]rune(s))
}
