package tests

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil { // caters for empty list
		return nil
	}

	current := l

	for i := 0; i < pos && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		return nil
	}
	return current
}
