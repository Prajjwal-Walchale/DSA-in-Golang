package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
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
	arr := []int{4, 3, 2, 1}
	list1 := &LinkedList{}
	for _, v := range arr {
		list1.insert(v)
	}

	list1.display()

}
