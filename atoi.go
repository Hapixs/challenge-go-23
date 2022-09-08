package piscine

func Atoi(s string) int {
	runes := atoiPopulateArray([]rune{}, s)
	total := 0
	totalNegative := false
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r >= 48 && r <= 57 {
			total = total*10 + int(r-48)
			continue
		} else if r == '+' && i >= 0 {
			if totalNegative && total == 0 {
				totalNegative = false
			} else if total > 0 {
				return 0
			}
		} else if r == '-' {
			if !totalNegative && total == 0 {
				totalNegative = true
			} else if total > 0 {
				return 0
			}
		} else {
			return 0
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
