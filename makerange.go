package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	tab := make([]int, max-min)
	for i := 0; i < len(tab); i++ {
		tab[i] = min + i
	}
	return tab
}
