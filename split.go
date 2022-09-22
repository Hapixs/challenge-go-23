package piscine

func Split(str, sep string) []string {
	for len(str) > 0 {
		i := Index(str, sep)
		if i < 0 {
			return []string{str}
		}
		return append([]string{str[:i]}, Split(str[i+len([]rune(sep)):], sep)...)
	}
	return []string{str}
}
