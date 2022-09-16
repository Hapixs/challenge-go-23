package piscine

func Join(s []string, sep string) string {
	if len(s) <= 0 {
		return ""
	}
	return s[0] + sep + Join(s[1:], sep)

}
