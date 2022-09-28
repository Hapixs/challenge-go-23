package piscine

func ListLast(l *List) interface{} {
	it := l.Head
	for it.Next != nil {
		it = it.Next
	}
	return it.Data
}
