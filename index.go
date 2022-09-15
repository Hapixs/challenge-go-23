package piscine

func Index(s, find string) int {
	rs := []rune(s)
	fs := []rune(find)

	for i, c := range s {
		if c == fs[0] {
			for y, fc := range find {
				if rs[i+y] != fc {
					break
				} else if y+1 == len(rs) {
					return i
				}
			}
		}
	}
	return -1
}
