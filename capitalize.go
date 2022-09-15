package piscine

func Capitalize(s string) string {
	r := []rune(s)
	str := ""
	for i := 0; i < len(r); i++ {
		c := r[i]
		if i == 0 {
			if IsLower(string(c)) {
				str += ToUpper(string(c))
			} else {
				str += string(c)
			}
			continue
		}
		if len(r)-1 >= i+1 && (c <= 47 || (c >= 58 && c <= 64) || (c >= 123 && c <= 126) || (c >= 91 && c <= 96)) && IsAlpha(string(r[i+1])) {
			str += string(c)
			if IsLower(string(r[i+1])) {
				str += ToUpper(string(r[i+1]))
			} else {
				str += string(r[i+1])
			}
			i++
			continue
		} else if IsUpper(string(c)) {
			str += ToLower(string(c))
			continue
		}
		str += string(c)
	}
	return str
}
