package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if len(r) < n || n < 0 {
		return '\x00'
	}
	return r[n-1]
}
