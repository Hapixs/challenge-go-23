package piscine

func Compare(a, b string) int {
	ar := []rune(a)
	br := []rune(b)

	if len(ar) > len(br) {
		return -1
	}
	if len(ar) < len(br) {
		return 1
	}
	if a == b {
		return 0
	}
	return 1
}
