package piscine

func SortIntegerTable(table []int) {
	m := make(map[int][]int)
	var maxValue int = table[0]
	var minValue int = table[0]
	for _, i := range table {
		m[i] = append(m[i], i)
		if i > maxValue {
			maxValue = i
		} else if i < minValue {
			minValue = i
		}
	}
	tb := []int{}
	for i := minValue; i <= maxValue; i++ {
		if val, ok := m[i]; ok {
			for _, v := range val {
				tb = append(tb, v)
			}
		}
	}
	copy(table, tb)
}
