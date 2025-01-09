package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	var current *NodeL = l
	var count int = 0
	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next
	}
	return nil
}
