package main

import (
	ls "github.com/prajjwal-w/data_structures_go/linkedList"
)

//var Node = linkedList.Node

// Delete a Node from last
func deleteANodefromLast(head *ls.Node) *ls.Node {
	if head == nil {
		return nil
	}

	temp := head
	for temp.Next.Next != nil {
		temp = temp.Next
	}

	temp.Next = nil

	return head

}

func main() {
	data := []interface{}{10, 20, 30, 40}
	linkedLst := ls.CreateLinkedList(data)

	//linkedlistTraversal
	ls.PrintLinkedList(linkedLst)

	//delete a node from last
	deleteANodefromLast(linkedLst)

	//linkedlistTraversal
	ls.PrintLinkedList(linkedLst)
}

// type Node struct {
// 	Data int
// 	Next *Node
// }

// func createALinkedList(data []int) *Node {
// 	var head *Node
// 	for _, value := range data {
// 		newNode := &Node{Data: value}
// 		if head == nil {
// 			head = newNode
// 		} else {
// 			curr := head
// 			for curr.Next != nil {
// 				curr = curr.Next
// 			}
// 			curr.Next = newNode
// 		}
// 	}
// 	return head
// }

// func PrintLinkedList(head *Node) {
// 	curr := head
// 	for curr != nil {
// 		fmt.Print(curr.Data, " -> ")
// 		curr = curr.Next
// 	}
// 	fmt.Print("nil\n")
// 	//fmt.Println()
// }
