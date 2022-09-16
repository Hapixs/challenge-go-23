package piscine

func LastRune(s string) rune {
	return NRune(s, len([]rune(s))-1)
}
