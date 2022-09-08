package piscine

func StrLen(s string) int {
	count := 0
	for _, c := range s {
		_ = c
		count++
	}
	return count
}
