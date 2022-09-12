package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 1; i < n; i++ {
		if (float64(n)/float64(i)) == float64(int64(n/i)) && n/i == i {
			return false
		}
	}
	return true
}
