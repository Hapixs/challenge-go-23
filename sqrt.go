package piscine

func Sqrt(n int) int {
	for i := 1; i < n; i++ {
		if n/i == i && (float64(n)/float64(i)) == float64(int64(n/i)) {
			return i
		}
	}
	return 0
}
