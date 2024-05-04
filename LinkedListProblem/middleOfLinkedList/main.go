package main

import (
	"fmt"
	"time"

	ll "github.com/prajjwal-w/data_structures_go/linkedList"
)

// MiddleOfTheList returns the middle node data
func MiddleOfTheList(head *ll.Node) interface{} {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Data
}

func main() {
	start := time.Now()
	data := []interface{}{20, 30, 40, 50, 60, 70}

	linkedList := ll.CreateLinkedList(data)

	//print ll
	ll.PrintLinkedList(linkedList)

	middleEle := MiddleOfTheList(linkedList)

	//print the middle of the linkedlist
	fmt.Println(middleEle)

	//check time of execution
	t := time.Since(start)
	fmt.Printf("execution time--->%v", t)

}
