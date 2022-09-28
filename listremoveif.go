package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	ListClear(l)
	for it != nil {
		if data_ref != it.Data {
			ListPushBack(l, it.Data)
		}
		it = it.Next
	}
}
