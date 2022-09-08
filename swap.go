package piscine

func Swap(a *int, b *int) {
	ta := *a
	*a = *b
	*b = ta
}
