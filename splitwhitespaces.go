package piscine

func SplitWhiteSpaces(str string) []string {
	for i, s := range str {
		if s == ' ' {
			return append([]string{str[:i]}, SplitWhiteSpaces(str[i+1:])...)
		}
	}
	return []string{}
}
