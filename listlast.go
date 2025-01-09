package piscine

func ListLast(l *List) interface{} {
	if ListSize(l) > 1 {
		var current *NodeL = l.Head
		var next *NodeL = current.Next

		for next != nil {
			current = next
			next = current.Next
		}
		return current.Data
	} else if ListSize(l) == 1 {
		return l.Head.Data
	} else {
		return nil
	}
}
