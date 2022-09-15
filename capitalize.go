package piscine

func Capitalize(s string) string {
	r := []rune(s)
	str := ""
	for i := 0; i < len(r); i++ {
		c := r[i]
		if i == 0 {
			if IsLower(string(c)) {
				str += ToUpper(string(r[i]))
				continue
			}
		}
		if len(r)-1 < i+1 && c < 47 && IsLower(string(r[i+1])) {
			str += string(c)
			str += ToUpper(string(r[i+1]))
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
