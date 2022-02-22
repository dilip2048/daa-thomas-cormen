/*
It is a linear data structure.

Search: O(n)
Insert: O(1)
Deletion: O(1)

*/

package main

import "fmt"

// Node is the placeholder for data and the next pointer
type Node struct {
	next *Node
	data interface{}
}

// This method will insert node to a linked list and makes it head
func insertNode(l *Node, data interface{}) *Node {
	//create a new node
	node := &Node{
		next: nil,
		data: data,
	}
	if l == nil {
		return node
	}
	node.next = l
	l = node
	return l
}

// This method will print the linked list
func (list *Node) printList() {
	temp := list
	for temp != nil {
		fmt.Printf("%v ", temp.data)
		temp = temp.next
	}
}

// This method will print the head of the linked list
func (list *Node) printHead() {
	fmt.Printf("%v", list.data)
}

// This method will search and remove the data
func (list *Node) remove(x int) {
	temp := list
	for temp.next.data != x {
		temp = temp.next
	}
	// the node to be deleted
	dNode := temp.next

	temp.next = temp.next.next

	// GC will take care of this node
	dNode.next = nil
}

func main() {
	var l *Node
	l = insertNode(l, 1)
	l = insertNode(l, 2)
	l = insertNode(l, 3)
	l = insertNode(l, 4)
	l.printHead()
	fmt.Println()
	l.printList()

	l.remove(3)
	fmt.Println()
	l.printList()
}
