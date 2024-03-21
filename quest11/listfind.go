package tests

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			// If the comparison indicates equality, return the address of the data interface
			return &current.Data
		}
		current = current.Next
	}
	// If no match is found, return nil
	return nil
}
