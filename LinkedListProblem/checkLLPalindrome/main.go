package main

import (
	"fmt"
	"time"

	ls "github.com/prajjwal-w/data_structures_go/linkedList"
	st "github.com/prajjwal-w/data_structures_go/stack"
)

// Brute force method
func isPalindrome(head *ls.Node) bool {
	//Initialized the curr pointer to the head of the linked list
	curr := head
	//create a stack to store the ele of linked list temporary
	stack := st.NewStack()

	//traverse the linked list and store the ele into stack
	for curr != nil {
		stack.Push(curr.Data)
		curr = curr.Next
	}

	//reset the curr pointer to the head of the list
	curr = head
	for curr != nil {
		ele, _ := stack.Pop()
		if curr.Data != ele {
			return false
		}

		curr = curr.Next
	}

	return true
}

// optimal apporach
func isPalindromeOptimal(head *ls.Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverse(slow)
	for slow != nil && (slow.Data == head.Data) {
		slow = slow.Next
		head = head.Next
	}

	return slow == nil
}

func reverse(head *ls.Node) *ls.Node {
	curr := head
	var prev *ls.Node
	for curr != nil {
		front := curr.Next
		curr.Next = prev
		prev = curr
		curr = front
	}

	return prev
}

func main() {
	t := time.Now()
	data := []interface{}{1, 2, 2, 1}

	linkedList := ls.CreateLinkedList(data)

	ls.PrintLinkedList(linkedList)

	//check palindrome bruteforce
	ans := isPalindrome(linkedList)
	//check palindrome optimal
	ansOP := isPalindromeOptimal(linkedList)

	//print true if palindrome otherwise false
	fmt.Println("Brute Force:")
	fmt.Println(ans)
	fmt.Println("Optimal:")
	fmt.Println(ansOP)
	fmt.Println(time.Since(t))
}
