package piscine

func BasicAtoi(s string) int {
	str := ""
	runes := []rune{}
	for _, st := range str {
		runes = append(runes, st)
	}

	for i, c := range s {
		if i > 0 && string(c) == string(runes[i-1]) {
			continue
		}
		str += string(c)
	}
	s = ""
	total := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j <= 9; j++ {
			if string(runes[i]) == string(rune(j+48)) {
				total = total*10 + j
			}
		}
	}
	return total
}
