package piscine

func Index(s, find string) int {
	rs := []rune(s)
	fs := []rune(find)

	if len(fs) > len(rs) {
		return -1
	}

	for i, c := range s {
		if c == fs[0] {
			for y, fc := range find {
				if rs[i+y] != fc {
					break
				}
				if y+1 == len(fs) {
					return i
				}
			}
		}
	}
	return -1
}
