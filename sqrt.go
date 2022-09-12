package piscine

func Sqrt(n int) int {
	for i := 1; i < n; i++ {
		if n/i == i && (n/i)%1.0 != 0 {
			return i
		}
	}
	return 0
}
