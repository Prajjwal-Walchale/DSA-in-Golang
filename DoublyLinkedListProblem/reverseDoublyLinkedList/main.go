package main

import (
	dls "github.com/prajjwal-w/data_structures_go/doublyLinkedList"
	st "github.com/prajjwal-w/data_structures_go/stack"
)

func ReverseDoublyLinkedListOptimal(head *dls.Node) *dls.Node {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var prev, front *dls.Node
	for curr != nil {
		front = curr.Next
		curr.Next = prev
		curr.Prev = front
		prev = curr
		curr = front

	}
	return prev
}

func ReverseDoublyLinkedListBruteForce(head *dls.Node) *dls.Node {
	curr := head
	//create stack
	stack := st.NewStack()
	for curr != nil {
		stack.Push(curr.Data)
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		ele, _ := stack.Pop()
		curr.Data = ele
		curr = curr.Next
	}
	return head
}

func main() {
	data := []interface{}{20, 30, 45, 50, 60, 70}

	//create doublyLinkedList
	doublyLS := dls.CreateDoublyLinkedList(data)

	//Print DLS
	dls.Display(doublyLS)

	// //Reverse DoublyLinkedList
	// reverseDLS := ReverseDoublyLinkedListBruteForce(doublyLS)

	// //Print reversed DLS
	// dls.Display(reverseDLS)

	//Reverse DLS Optimal
	revDls := ReverseDoublyLinkedListOptimal(doublyLS)

	//Print reversed DLS
	dls.Display(revDls)

}
