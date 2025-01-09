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
	NodeL := &NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = NodeL
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = NodeL
	}
	return
}
