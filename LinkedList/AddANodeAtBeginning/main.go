package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBeg(data int) {
	newNode := &Node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) display() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Print("nil")
}

func main() {
	l1 := &LinkedList{}
	l1.insertAtBeg(2)
	l1.insertAtBeg(3)
	l1.insertAtBeg(5)
	l1.display()
}
