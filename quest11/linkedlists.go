package tests

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
func main() {

	var link *node
	link = insert(link, 1)
	link = insert(link, 2)
	link = insert(link, 3)
	link = insert(link, 4)
	link = insert(link, 5)

	PrintList(link)

}
