package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	current := l.Head
	for {
		if current == nil {
			current = l.Tail
			break
		} else {
			current = current.Next
		}
	}
	l.Tail = &NodeL{data, &NodeL{}}
}
