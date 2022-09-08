package piscine

func SortIntegerTable(table []int) {
	m := make(map[int]int)
	var maxValue int = table[0]
	var minValue int = table[0]
	for _, i := range table {
		m[i] = i
		if i > maxValue {
			maxValue = i
		}
		if i < minValue {
			minValue = i
		}
	}
	tb := []int{}
	for i := minValue; i <= maxValue; i++ {
		if val, ok := m[i]; ok {
			tb = append(tb, val)
		}
	}
	copy(table, tb)
}
