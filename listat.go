package piscine

func ListAt(node *NodeL, at int) *NodeL {
	index := 0
	it := node
	for index < at {
		if it.Next == nil {
			return nil
		}
		it = it.Next
		index++
	}
	return it
}
