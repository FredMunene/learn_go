package tests

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head} // create a new node

	if l.Head == nil { // if list is emmpty, both Head and Tail point to new node
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head // new node points to current head
	l.Head = newNode      // Head updated to be new
}
