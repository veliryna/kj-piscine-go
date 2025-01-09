package piscine

func ListPushFront(l *List, data interface{}) {
	NodeL := &NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = NodeL
	} else {
		NodeL.Next = l.Head
		l.Head = NodeL
	}
	return
}
