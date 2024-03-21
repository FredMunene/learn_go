package tests

func SortListInsert(l *NodeL, data_ref int) *NodeL {
	newNode := &NodeL{Data: data_ref}
	// If the list is empty or the new data is smaller than the head, insert at the beginning
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return l
}
