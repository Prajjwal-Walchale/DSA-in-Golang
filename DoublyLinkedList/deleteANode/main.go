package main

import (
	dls "github.com/Prajjwal-Walchale/DSA-in-Golang/doublyLinkedList"
)

func deleteANodeFromLast(head *dls.Node) *dls.Node {
	if head == nil || head.Next == nil {
		return nil
	}
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next.Prev = nil
	curr.Next = nil

	return head
}

func main() {
	data := []interface{}{20, 40, 55, 23, 56, 78, 90}

	//create dls
	doublyLinkedList := dls.CreateDoublyLinkedList(data)
	//displaylist
	dls.Display(doublyLinkedList)
	//delete last node
	updatedList := deleteANodeFromLast(doublyLinkedList)

	dls.Display(updatedList)

}
