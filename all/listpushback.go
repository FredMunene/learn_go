package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// that inserts a new element NodeL at the end of the list l while using the structure List.
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode
	l.Tail = newNode
}
