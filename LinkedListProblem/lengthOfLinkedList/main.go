package main

import (
	"fmt"

	ls "github.com/prajjwal-w/data_structures_go/linkedList"
)

func CountNodeOfLinkedList(head *ls.Node) int {
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func main() {
	data := []interface{}{12, 24, 16, 15, 19}

	//create linkedList
	linkedList := ls.CreateLinkedList(data)

	//count nodes
	count := CountNodeOfLinkedList(linkedList)

	fmt.Println(count)
}
