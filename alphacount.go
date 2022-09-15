package piscine

func AlphaCount(s string) int {
	i := 0
	for _, c := range s {
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			i += 1
		}
	}
	return i
}
