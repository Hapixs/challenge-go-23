package piscine

func NRune(s string, n int) rune {
	return map[bool]rune{true: []rune(s)[n], false: '\x00'}[len([]rune(s)) < n || n <= 0]
}
