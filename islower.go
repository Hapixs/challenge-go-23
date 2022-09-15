package piscine

func IsLower(s string) bool {
	for _, c := range s {
		if !(c >= 97 && c <= 122) {
			return false
		}
	}
	return true
}
