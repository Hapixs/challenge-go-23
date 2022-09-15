package piscine

func AtoiBase(s string, base string) int {
	for x, c := range base {
		for y, c1 := range base {
			if len([]rune(base)) <= 1 || (c == c1 && x != y) || c1 == '-' || c1 == '+' {
				return 0
			}
		}
	}

	nbase := len([]rune(base))
	s = StrRev(s)
	toConvert := 0
	for i, c := range s {
		for y, b := range base {
			if c == b {
				toConvert += y * pow(nbase, nbase, i)
			}
		}
	}
	return toConvert
}

func pow(n, nb, rpower int) int {
	if n == nb && rpower < 1 {
		return 1
	}
	if rpower <= 1 {
		return nb
	}
	return recPower(n, nb*n, rpower-1)
}
