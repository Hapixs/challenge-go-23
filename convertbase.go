package piscine

func ConvertBase(nbr, fromBase, toBase string) string {
	number := AtoiBase(nbr, fromBase)
	return ToBase(number, toBase)
}

func ToBase(nbr int, base string) string {
	return ToBaseRec(nbr, []rune(base))
}

func ToBaseRec(nbr int, baseTo []rune) string {
	if nbr < len(baseTo)-1 {
		return string(baseTo[nbr])
	}
	return ToBaseRec(nbr/(len([]rune(baseTo))-1), baseTo)
}
