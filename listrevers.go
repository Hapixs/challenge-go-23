package piscine

func ListReverse(l *List) *List {
	list := &List{}
	it := l.Head
	if it == nil {
		return nil
	}
	for it.Next != nil {
		ListPushFront(list, it.Data)
	}
	return l
}
