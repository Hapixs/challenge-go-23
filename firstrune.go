package piscine

func FirstRune(s string) rune {
	if len([]rune(s)) <= 0 {
		return '\x00'
	}
	return []rune(s)[0]
}
