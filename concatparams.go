package piscine

func ConcatParams(str []string) string {
	s := ""
	for _, st := range str {
		s += st + "\n"
	}
	return s
}
