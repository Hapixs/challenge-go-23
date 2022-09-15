package piscine

func AtoiBase(s string, base string) int {
	for x, c := range base {
		for y, c1 := range base {
			if len([]rune(base)) <= 1 || (c == c1 && x != y) || c1 == '-' || c1 == '+' {
				return 0
			}
		}
	}
	fnl := 0
	for _, c := range s {
		for i, b := range base {
			if c == b {
				fnl = (10 * fnl) + i
			}
		}
	}
	return fnl
}
