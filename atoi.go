package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	str := ""
	runes := atoiPopulateArray([]rune{}, s)
	patchZero := false
	patchPos := false
	for i, c := range runes {
		if !patchZero && i >= 0 && c != '0' {
			patchZero = true
		}
		if !patchZero && c == '0' {
			continue
		}
		if !patchPos && i > 0 && c == '+' {
			patchPos = true
		}
		if patchPos && c == '+' {
			return 0
		}
		str += string(c)
	}
	runes = atoiPopulateArray([]rune{}, str)
	total := 0
	totalNegative := false
	for i := 0; i < len(runes); i++ {
		for j := 0; j <= 10; j++ {
			if j == 10 {
				if runes[i] == ' ' || runes[i] == '+' {
					break
				}
				if runes[i] == '-' && i == 0 {
					if totalNegative {
						return 0
					}
					totalNegative = true
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
	if totalNegative {
		total *= -1
	}
	return total
}

func atoiPopulateArray(runes []rune, str string) []rune {
	for _, st := range str {
		runes = append(runes, st)
	}
	return runes
}
