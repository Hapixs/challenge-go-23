package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	it := n1
	for it.Next != nil {
		it = it.Next
	}
	it.Next = n2
	ListSort(n1)
	return n1
}
