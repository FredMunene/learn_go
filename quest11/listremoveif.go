package tests

func ListRemoveIf(l *List, data_ref interface{}) {
	prev := &NodeL{Next: l.Head}
	current := l.Head
	for current != nil {
		if current.Data == data_ref {
			if current == l.Head {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			if current == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
