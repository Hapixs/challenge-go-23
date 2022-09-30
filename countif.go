package piscine

func CountIf(f func(string) bool, a []string) int {
	am := 0
	for _, s := range a {
		am = map[bool]int{true: am + 1, false: am}[f(s)]
	}
	return am
}
