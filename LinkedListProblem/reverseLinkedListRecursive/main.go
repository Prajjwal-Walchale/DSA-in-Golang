package main

import (
	ll "github.com/prajjwal-w/data_structures_go/linkedList"
)

func reverseLinkedList(head *ll.Node) *ll.Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseLinkedList(head.Next)

	front := head.Next

	front.Next = head

	head.Next = nil

	return newHead

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
