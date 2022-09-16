package piscine

func Join(s []string, sep string) string {
	if len(s) <= 0 {
		return ""
	}
	return s[0] + map[bool]string{true: sep, false: ""}[len(s) > 1] + Join(s[1:], sep)
}
