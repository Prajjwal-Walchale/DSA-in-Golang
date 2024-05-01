package main

import (
	dls "github.com/Prajjwal-Walchale/DSA-in-Golang/doublyLinkedList"
	st "github.com/Prajjwal-Walchale/DSA-in-Golang/stack"
)

func ReverseDoublyLinkedListBruteForce(head *dls.Node) *dls.Node {
	curr := head

	//create stack
	stack := st.NewStack{}
	for curr != nil {

	}
}

func main() {
	data := []interface{}{20, 30, 45, 50, 60, 70}

	//create doublyLinkedList
	doublyLS := dls.CreateDoublyLinkedList(data)

	//Print DLS
	dls.Display(doublyLS)

	//Reverse DoublyLinkedList
	reverseDLS := ReverseDoublyLinkedListBruteForce(doublyLS)

}
