package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 1000000 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return factorise(nb, 1, 1) * nb
}

func factorise(n int, index int, fact int) int {
	if index == n-1 {
		return fact * index
	}
	return factorise(n, index+1, fact*index)
}
