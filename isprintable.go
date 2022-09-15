package piscine

func IsPrintable(s string) bool {
	for _, c := range s {
		if c <= 31 {
			return false
		}
	}
	return true
}
