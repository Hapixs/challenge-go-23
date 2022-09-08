package piscine

func StrRev(s string) string {
	runes := []rune{}
	for _, st := range s {
		runes = append(runes, st)
	}
	s = ""
	for i := len(runes); i >= 0; i-- {
		s += string(runes[i])
	}
	return s
}
