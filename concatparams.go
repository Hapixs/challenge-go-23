package piscine

func ConcatParams(str []string) string {
	s := ""
	for i, st := range str {
		s += st
		if i < len(str)-1 {
			s += "\n"
		}
	}
	return s
}
