package main

import (
	"fmt"

	dls "github.com/Prajjwal-Walchale/DSA-in-Golang/doublyLinkedList"
)

func insertAtTheEnd(head *dls.Node, ele int) *dls.Node {
	newNode := &dls.Node{Data: ele}
	if head == nil {
		head = newNode
		return head
	}

	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
	newNode.Prev = curr

	return head

}

func main() {
	data := []interface{}{20, 40, 55, 23, 56, 78, 90}

	//create doubly linked list
	doubleLinkedList := dls.CreateDoublyLinkedList(data)

	//print doubly linked list
	dls.Display(doubleLinkedList)

	//Insert at the end
	updatedDLS := insertAtTheEnd(doubleLinkedList, 100)

	//print DLS2
	fmt.Println()
	dls.Display(updatedDLS)
}
