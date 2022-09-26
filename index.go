package piscine

func Index(s, find string) int {
	SizeS := StrLen(s)
	SizeF := StrLen(find)
	for i := 0; i <= SizeS-SizeF; i++ {
		if s[i:i+SizeF] == find {
			return i
		}
	}
	return -1
}
