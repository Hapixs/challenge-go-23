package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a // echange des valeurs
}
