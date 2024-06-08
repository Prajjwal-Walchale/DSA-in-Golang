package main

import (
	ll "github.com/prajjwal-w/data_structures_go/linkedList"
)

func seggregateOddandEven(head *ll.Node) *ll.Node {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead := head
	evenHead := head.Next
	currOdd := oddHead
	currEven := evenHead
	for currEven != nil && currEven.Next != nil {
		currOdd.Next = currEven.Next
		currOdd = currOdd.Next
		currEven.Next = currOdd.Next
		currEven = currEven.Next

	}
	currOdd.Next = evenHead

	return oddHead
}

func main() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	linkedL := ll.CreateLinkedList(data)
	//print linkedlist
	ll.PrintLinkedList(linkedL)

	segList := seggregateOddandEven(linkedL)

	ll.PrintLinkedList(segList)
}
