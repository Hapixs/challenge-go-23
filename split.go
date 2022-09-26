package piscine

func Split(str, sep string) []string {
	if len(str) > 0 {
		i := Index(str, sep)
		if i >= 0 {
			return append([]string{str[:i]}, Split(str[i+len([]rune(sep)):], sep)...)
		}
	}
	return []string{str}
}
