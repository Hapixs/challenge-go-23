package piscine

func ToLower(s string) string {
	if len(s) > 0 {
		r := rune(s[0])
		return string(rune(int(r)+map[bool]int{true: 32, false: 0}[r >= 65 && r <= 90])) + ToUpper(s[1:])
	}
	return s
}
