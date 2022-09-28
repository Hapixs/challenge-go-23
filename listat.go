package piscine

func ListAt(l *List, at int) *NodeL {
	if l.Head == nil || ListSize(l) < at {
		return nil
	}
	index := 0
	it := l.Head
	for index < at {
		it = it.Next
		index++
	}
	return it
}
