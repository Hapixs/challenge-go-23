package piscine

func SplitWhiteSpaces(str string) []string {
	for i, s := range str {
		if s == ' ' || s == '\n' || s == '\t' && len(str) > 0 {
			if len(str[:i]) > 0 {
				return append([]string{str[:i]}, SplitWhiteSpaces(str[i+1:])...)
			} else {
				return append([]string{}, SplitWhiteSpaces(str[i+1:])...)
			}
		}
	}
	if len(str) > 0 {
		return []string{str}
	}
	return []string{}
}
