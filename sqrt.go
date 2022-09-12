package piscine

func Sqrt(n int) int {
	if n == 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		if n/i == i && (n/i)%1.0 != 0 {
			return i
		}
	}
	return 0
}
