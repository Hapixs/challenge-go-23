package piscine

func IsUpper(s string) bool {
	for _, c := range s {
		if !(c >= 65 && c <= 90) {
			return false
		}
	}
	return true
}
