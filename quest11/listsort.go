package tests

func ListSort(l *NodeL) *NodeL {
	if l == nil || l.Next == nil {
		return l
	}

	// Split the list into two halves
	var (
		fast = l.Next
		slow = l
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := slow.Next
	slow.Next = nil

	// Recursively sort the two halves
	first := ListSort(l)
	second = ListSort(second)

	// Merge the sorted halves
	return merge(first, second)
}

func merge(l1, l2 *NodeL) *NodeL {
	dummy := &NodeL{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Append remaining nodes
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}
