package piscine

func Capitalize(s string) string {
	r := []rune(s)
	str := ""
	for i, c := range r {
		if i == 0 {
			if IsLower(string(c)) {
				str += ToUpper(string(r[i]))
				continue
			}
		}
		if len(r)-1 < i+1 && c < 47 && IsLower(string(r[i+1])) {
			str += ToUpper(string(r[i+1]))
			continue
		}
		str += string(c)
	}
	return str
}
