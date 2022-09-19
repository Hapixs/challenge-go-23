package piscine

func ToBaseRec(nbr int, baseTo []rune, size int) string {
	if nbr < size {
		return string(baseTo[nbr])
	}
	return ToBaseRec(nbr/size, baseTo, size) + string(baseTo[nbr%size])
}

func ToBase(nbr int, base string) string {
	return ToBaseRec(nbr, []rune(base), len([]rune(base)))
}

func ConvertBase(nbr, fromBase, toBase string) string {
	number := AtoiBase(nbr, fromBase)
	return ToBase(number, toBase)
}
