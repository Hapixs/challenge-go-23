package piscine

func IsAlpha(s string) bool {
	for _, c := range s {
		if !(c >= 97 && c <= 122) && !(c >= 65 && c <= 90) && !(c >= 58 && c <= 57) {
			return false
		}
	}
	return true
}
