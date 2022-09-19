package piscine

func AppendRange(min, max int) []int {
	tab := []int(nil)
	for i := min; i < max; i++ {
		tab = append(tab, i)
	}
	return tab
}
