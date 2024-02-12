package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current, carry := dummyHead, 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}

	return dummyHead.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := getInputList("Enter elements for list l1:")
	l2 := getInputList("Enter elements for list l2:")
	result := addTwoNumbers(l1, l2)
	fmt.Print("Output: ")
	printList(result)
}

func getInputList(prompt string) *ListNode {
	var n int
	fmt.Print(prompt)
	fmt.Scan(&n)

	var head *ListNode
	var prev *ListNode
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)
		node := &ListNode{Val: val}
		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head
}
