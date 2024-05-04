package main

import (
	ll "github.com/prajjwal-w/data_structures_go/linkedList"
)

func reverseLinkedList(head *ll.Node) *ll.Node {
	curr := head
	var prev *ll.Node
	for curr != nil {
		front := curr.Next
		curr.Next = prev
		prev = curr
		curr = front
	}

	return prev
}

func main() {
	data := []interface{}{20, 30, 40, 60, 70, 90}

	//create linkedlist
	linkedL := ll.CreateLinkedList(data)
	//print linkedlist
	ll.PrintLinkedList(linkedL)

	reverseLL := reverseLinkedList(linkedL)

	ll.PrintLinkedList(reverseLL)
}
