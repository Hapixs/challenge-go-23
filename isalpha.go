package piscine

func IsAlpha(s string) bool {
	for _, c := range s {
		if !(c >= 97 && c <= 122) && !(c >= 65 && c <= 90) {
			return false
		}
	}
	return true
}
