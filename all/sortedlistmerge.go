package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	dummy := &NodeI{}
	current := dummy

	// Merge the two lists
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Append remaining nodes
	if n1 != nil {
		current.Next = n1
	} else if n2 != nil {
		current.Next = n2
	}

	return dummy.Next
}

func listPushBack(l *NodeI, data int) *NodeI {
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
