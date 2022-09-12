package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	i := 1
	for i = 2; i < n; i++ {
		if (float64(n) / float64(i)) == float64(int64(n/i)) {
			return false
		}
	}
	return true
}
