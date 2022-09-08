package piscine

func BasicAtoi2(s string) int {
	runes := []rune{}
	for _, st := range s {
		runes = append(runes, st)
	}
	total := 0
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r >= 48 && r <= 57 {
			total = total*10 + int(r-48)
			continue
		} else {
			return 0
		}
	}
	return total
}
