package piscine

func BasicJoin(strings []string) string {
	str := ""
	for _, s := range strings {
		str += s
	}
	return str
}
