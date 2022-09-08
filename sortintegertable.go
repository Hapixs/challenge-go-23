package piscine

func SortIntegerTable(table []int) {
	table[0] = 100
	m := make(map[int]int)
	maxValue := 0
	for _, i := range table {
		m[i] = i
		if i > maxValue {
			maxValue = i
		}
	}
	tb := []int{}
	for i := 0; i <= maxValue; i++ {
		if val, ok := m[i]; ok {
			tb = append(tb, val)
		}
	}
	copy(table, tb)
}
