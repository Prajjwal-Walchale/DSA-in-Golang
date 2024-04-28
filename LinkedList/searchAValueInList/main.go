package main

import (
	"fmt"

	ls "github.com/Prajjwal-Walchale/DSA-in-Golang/linkedList"
)

func FindTheElememtInTheList(head *ls.Node, ele int) int {
	curr := head
	for curr != nil {
		if curr.Data == ele {
			return 1
		}
		curr = curr.Next
	}

	return 0
}

func main() {
	fmt.Println("statrted")
	data := []int{1, 2, 3, -1, -2, 5}

	//create linkedlist from the given data
	linkedList := ls.CreateLinkedList(data)

	//element to find
	ele := -1
	//search the element in the linkedlist
	ans := FindTheElememtInTheList(linkedList, ele)
	fmt.Println("If the element is present print 1 or print 0")
	fmt.Println(ans)

}
