package piscine

func SplitWhiteSpaces(str string) []string {
	for i, s := range str {
		if s == ' ' || s == '\n' || s == '\t' {
			if len(str[:i]) > 0 {
				return append([]string{str[:i]}, SplitWhiteSpaces(str[i+1:])...)
			}
		}
	}
	return []string{str}
}
