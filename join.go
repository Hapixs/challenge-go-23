package piscine

func Join(s []string, sep string) string {
	if len(s) <= 0 {
		return ""
	}
	str := s[0]
	for _, st := range s[0:] {
		str += sep + st
	}
	return str
}
