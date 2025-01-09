package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	node2 := n2
	for node2 != nil {
		n1 = SortListInsert(n1, node2.Data)
		node2 = node2.Next
	}
	return n1
}
