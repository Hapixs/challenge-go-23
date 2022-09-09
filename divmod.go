package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div, *mod = a/b, a%b // On affect les nouvelle valeurs au variable pointer par les pointeurs
}
