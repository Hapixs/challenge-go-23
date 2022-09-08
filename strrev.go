package piscine

func StrRev(s string) {
	runes := []rune{}

	for _, st := range s {
		runes = append(runes, st)
	}
	for i := len(runes); i >= 0; i-- {
		print(runes[i])
	}
}
