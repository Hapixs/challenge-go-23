package piscine

func SortListInsert(nodeI *NodeI, ref int) {
	it := nodeI
	for it.Next != nil {
		it = it.Next
	}
	it.Next = &NodeI{Data: ref}
	ListSort(nodeI)
}
