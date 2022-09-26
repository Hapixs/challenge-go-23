package piscine

func SortWordArr(s []string) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s) && rune(s[j-1][0]) < rune(s[j][0]); j++ {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
