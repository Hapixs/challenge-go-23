package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}
	return factorise(nb) * nb
}

func factorise(n int) int {
	fact := 1
	for i := 1; i < n; i++ {
		fact = fact * i
	}
	return fact
}
