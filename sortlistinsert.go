package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := l
	counter := 0
	var index int
	var prev *NodeI
	for node != nil {
		if node.Data >= data_ref && counter == 0 {
			index = 0
			break
		}
		if node.Data >= data_ref && prev.Data <= data_ref && prev != nil {
			index = counter
			break
		}
		counter++
		prev = node
		node = node.Next
	}
	if node == nil {
		return pushBack(l, data_ref)
	}
	return InsertNodeAtIthIndex(l, index, data_ref)
}

func NewNode(value int, next *NodeI) *NodeI {
	var n NodeI
	n.Data = value
	n.Next = next
	return &n
}

func InsertNodeAtIthIndex(head *NodeI, index, data int) *NodeI {
	if head == nil {
		head = NewNode(data, nil)
		return head
	}
	if index == 0 {
		newNode := NewNode(data, nil)
		newNode.Next = head
		head = newNode
		return head
	}
	i := 0
	temp := head
	preNode := temp
	for temp != nil {
		if i == index {
			newNode := NewNode(data, nil)
			preNode.Next = newNode
			newNode.Next = temp
			break
		}
		i++
		preNode = temp
		temp = temp.Next
	}
	return head
}

func pushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
