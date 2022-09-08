package piscine

func UltimateDivMod(a *int, b *int) {
	ta := *a
	tb := *b

	*a = ta / tb
	*b = ta % tb
}
