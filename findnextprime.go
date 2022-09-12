package piscine

func FindNextPrime(n int) int {
	n += 1
	for !IsPrime(n) {
		n += 1
	}
	return n
}
