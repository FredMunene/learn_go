package piscine

import "fmt"

type node struct {
	data      int
	next_node *node
}

func insert(head *node, data int) *node {
	n := &node{data: data}
	if head == nil {
		return n
	} else {
		n.next_node = head
		return n
	}
}

func PrintList(head *node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next_node
	}
	fmt.Println(nil)
}
