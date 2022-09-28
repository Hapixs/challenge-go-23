package piscine

func ListReverse(l *List) {
	it := l.Head
	ListClear(l)
	for it != nil {
		ListPushFront(l, it.Data)
		it = it.Next
	}
}
