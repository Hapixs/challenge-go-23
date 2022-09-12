package piscine

func RecursivePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return recPower(nb, nb, power)
}

func recPower(n, nb, rpower int) int {
	if rpower <= 0 {
		return nb
	}
	return recPower(n, nb*n, rpower-1)
}
