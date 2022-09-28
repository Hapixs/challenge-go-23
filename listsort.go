package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	i := l
	for i != nil {
		y := i.Next
		for y != nil {
			if i.Data > y.Data {
				i.Data, y.Data = y.Data, i.Data
			}
			y = y.Next
		}
		i = i.Next
	}
	return l
}
