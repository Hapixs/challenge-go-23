package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	str := ""
	runes := atoiPopulateArray([]rune{}, s)
	patchZero := false
	for i, c := range runes {
		if !patchZero && i >= 0 && c != '0' {
			patchZero = true
		}
		if !patchZero && c == '0' {
			continue
		}
		str += string(c)
	}
	runes = atoiPopulateArray([]rune{}, str)

	total := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j <= 10; j++ {
			if j == 10 {
				if runes[i] == ' ' {
					break
				}
				if runes[i] == '-' {
					if total > 0 {
						total *= -1
					}
					break
				}
				if runes[i] == '+' {
					if total < 0 {
						total *= -1
					}
					break
				}
				return 0
			}
			if string(runes[i]) == string(rune(j+48)) {
				total = total*10 + j
				break
			}
		}
	}
	return total
}

func atoiPopulateArray(runes []rune, str string) []rune {
	for _, st := range str {
		runes = append(runes, st)
	}
	return runes
}
