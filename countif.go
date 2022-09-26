package piscine

func CountIf(f func(string) bool, a []string) int {
	am := 0
	for _, s := range a {
		if f(s) {
			am++
		}
	}
	return am
}
