package piscine

func ToLower(s string) string {
	str := ""
	for _, c := range s {
		if c >= 65 && c <= 90 {
			str += string(c + 32)
		} else {
			str += string(c)
		}
	}
	return str
}
