package piscine

func StrRev(s string) string {
	str := ""
	for _, c := range s {
		str = string(c) + str
	}
	return str
}
