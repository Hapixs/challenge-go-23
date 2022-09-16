package piscine

func AlphaCount(s string) int {
	if StrLen(s) < 1 {
		return 0
	}
	return map[bool]int{true: map[bool]int{true: 1 + AlphaCount(string(s[1:])), false: 0}[isLetter([]rune(s)[0])], false: 0}[StrLen(s) > 0]
}

func isLetter(c rune) bool {
	return (c >= 97 && c <= 122) || (c >= 65 && c <= 90)
}
