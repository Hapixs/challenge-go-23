package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[j-1], a[j]) == a[j] {
				return false
			}
		}
	}
	return true
}
