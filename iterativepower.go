package piscine

func IterativePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	fnl := nb
	for i := 0; i < power; i++ {
		nb *= nb
	}
	return fnl
}
