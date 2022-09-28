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
	l.Head.Next = l.Tail
	current := l.Head
	for {
		if current.Next == nil {
			current.Next = l.Tail
			break
		} else {
			current = current.Next
		}
	}
	l.Tail = &NodeL{data, &NodeL{}}
}
