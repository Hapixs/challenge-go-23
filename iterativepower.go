package piscine

func IterativePower(nb, power int) int {
	fnl := nb
	for i := 1; i < power; i++ {
		fnl *= nb
	}

	return fnl
}
