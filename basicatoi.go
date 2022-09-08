package piscine

import "strings"

func BasicAtoi(s string) int {
	str := ""

	for i, c := range s {
		if i > 0 && string(c) == strings.Split(s, "")[i-1] {
			continue
		}
		str += string(c)
	}
	runes := []rune{}
	for _, st := range str {
		runes = append(runes, st)
	}
	s = ""
	total := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j <= 9; j++ {
			if strings.Split(str, "")[i] == string(rune(j+48)) {
				total = total*10 + j
			}
		}
	}
	return total
}
