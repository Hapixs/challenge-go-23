package piscine

func IterativePower(nb, power int) int {
	if nb < 0 {
		return 1
	}
	fnl := nb
	for i := 1; i < power; i++ {
		fnl *= nb
	}

	return fnl
}
