package piscine

func Join(s []string, sep string) string {
	if len(s) <= 0 {
		return ""
	}
	str := s[0]
	for i := 1; i < len(s); i++ {
		str += sep + s[i]
	}
	return str
}
