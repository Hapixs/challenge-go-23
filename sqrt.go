package piscine

func Sqrt(n int) int {
	if n <= 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if n/i == i {
			return i
		}
	}
	return 0
}
