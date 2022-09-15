package piscine

func ToUpper(s string) string {
	str := ""
	for _, c := range s {
		if c >= 97 && c <= 122 {
			str += string(c - 32)
		} else {
			str += string(c)
		}
	}
	return str
}
