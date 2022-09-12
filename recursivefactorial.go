package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 1000000 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return factoriseRec(nb, 1, 1)
}

func factoriseRec(n int, index int, fact int) int {
	if index == n {
		return fact * index
	}
	return factoriseRec(n, index+1, fact*index)
}
