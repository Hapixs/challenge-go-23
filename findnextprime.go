package piscine

func FindNextPrime(n int) int {
	for !IsPrime(n) {
		n += 1
	}
	return n
}
