package piscine

func LastRune(s string) rune {
	if len([]rune(s)) <= 0 {
		return '\x00'
	}
	return []rune(s)[len([]rune(s))-1]
}
