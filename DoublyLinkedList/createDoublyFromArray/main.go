package main

import (
	"fmt"

	dls "github.com/Prajjwal-Walchale/DSA-in-Golang/doublyLinkedList"
)

func check(head *dls.Node, ele int) {
	curr := head
	for curr != nil {
		if curr.Data == ele {
			fmt.Println(curr.Prev.Data, " <-> ", curr.Data, " <-> ", curr.Next.Data)
		}
		curr = curr.Next
	}
}

func main() {
	data := []interface{}{3, 5, 7, 8, 9}

	//create doubly linked list
	doublyLinkedList := dls.CreateDoublyLinkedList(data)

	//print doublyLinkedList
	dls.Display(doublyLinkedList)
	fmt.Println()
	//check func
	check(doublyLinkedList, 7)
}
