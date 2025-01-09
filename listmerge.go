package piscine

func ListMerge(l1 *List, l2 *List) {
	node := l2.Head
	for node != nil {
		ListPushBack(l1, node.Data)
		node = node.Next
	}
}
