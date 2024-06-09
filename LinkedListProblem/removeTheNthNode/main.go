package main

import (
	ll "github.com/prajjwal-w/data_structures_go/linkedList"
)

func removeNthNode(head *ll.Node, n int) *ll.Node {
	first := head
	sec := head

	for i := 0; i < n; i++ {
		first = first.Next
	}

	if first == nil {
		return head.Next
	}

	for first.Next != nil {
		first = first.Next
		sec = sec.Next
	}

	sec.Next = sec.Next.Next

	return head
}

func main() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	linkedlist := ll.CreateLinkedList(data)

	//printlinked list
	ll.PrintLinkedList(linkedlist)

	//node from last
	n := 2
	//remove the nth node from the last
	list := removeNthNode(linkedlist, n)

	//Printlinkedlist
	ll.PrintLinkedList(list)

}
