package piscine

func ListSize(l *List) int {
	var temp *NodeL = l.Head
	var count int = 0
	for temp != nil {
		temp = temp.Next
		count++
	}
	return count
}
