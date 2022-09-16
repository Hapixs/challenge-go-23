package piscine

func Join(s []string, sep string) string {
	return map[bool]string{true: s[0] + map[bool]string{true: sep, false: ""}[len(s) > 1] + Join(s[1:], sep), false: ""}[len(s) >= 1]
}
