package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && f(a[j-1], a[j]) >= 1; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}
