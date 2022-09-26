package piscine

func SortWordArr(s []string) {
	ru := []rune{}
	for _, str := range s {
		ru = append(ru, rune(str[0]))
	}
	for i := 0; i < len(ru); i++ {
		for j := i + 1; j < len(ru) && rune(s[j-1][0]) > rune(s[j][0]); j++ {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
