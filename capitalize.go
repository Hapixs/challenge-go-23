package piscine

func Capitalize(s string) string {
	r := []rune(s)
	for i, c := range r {
		if i == 0 && IsAlpha(string(c)) {
			r[i] = rune(ToUpper(string(c))[0])
		} else if IsAlpha(string(c)) && !IsAlpha(string(r[i-1])) {
			r[i] = rune(ToUpper(string(c))[0])
		} else if !IsLower(string(c)) {
			r[i] = rune(ToLower(string(c))[0])
		}
	}
	return string(r)
}
