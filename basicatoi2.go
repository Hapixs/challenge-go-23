package piscine

func BasicAtoi2(s string) int {
	if s == "" {
		return 0
	}
	str := ""
	runes := []rune{}
	for _, st := range s {
		runes = append(runes, st)
	}

	for i, c := range runes {
		if i > 0 && c == runes[i-1] {
			continue
		}
		str += string(c)
	}
	s = ""
	total := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j <= 10; j++ {
			if i == 10 {
				return 0
			}
			if string(runes[i]) == string(rune(j+48)) {
				total = total*10 + j
			}
		}
	}
	return total
}
