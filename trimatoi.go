package piscine

func TrimAtoi(s string) int {
	negative := false
	total := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			total = total*10 + int(c-48)
			continue
		} else if (c == '-' || c == '+') && total <= 0 {
			negative = c == '-'
			continue
		}
	}

	return map[bool]int{true: total * -1, false: total}[negative] // Ici on retourn le chiffre en positif si (negative == false) ou en negatif si (negative == true)
}
