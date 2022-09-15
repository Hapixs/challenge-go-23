package piscine

func Base(nbr int, base string) {
	nbase := len([]rune(base))
	str := ""
	for nbr/nbase > 0 {
		str = string(base[nbr%nbase]) + str
		nbr = nbr / nbase
	}
	println(string(base[nbr]) + str)
}
