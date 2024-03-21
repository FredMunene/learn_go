package tests

func ListSize(l *List) int { // counts number of elements in a linked list>>no of nodes??
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
