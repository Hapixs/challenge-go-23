package piscine

func IterativePower(nb, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 && nb < 0 {
		return 0
	}
	fnl := nb
	for i := 1; i < power; i++ {
		fnl *= nb
	}
	return fnl
}
