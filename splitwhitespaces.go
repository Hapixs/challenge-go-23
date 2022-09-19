package piscine

func SplitWhiteSpaces(str string) []string {
	return addInTabRecusrive(str, []string{})
}

func addInTabRecusrive(str string, tab []string) []string {
	for i, st := range str {
		if st == ' ' {
			tab = append(tab, str[:i-1])
			return append(tab, addInTabRecusrive(str[i-1:i+1], tab)...)
		}
	}
	return tab
}
