package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := []bool{}
	for _, i := range a {
		tab = append(tab, f(i))
	}
	return tab
}
