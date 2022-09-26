package piscine

func ToUpper(s string) string {
	if len(s) > 0 {
		r := rune(s[0])
		return string(rune(int(r)-map[bool]int{true: 32, false: 0}[r >= 97 && r <= 122])) + ToUpper(s[1:])
	}
	return s
}
