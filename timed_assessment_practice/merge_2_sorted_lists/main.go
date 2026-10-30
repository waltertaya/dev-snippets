package main

import "fmt"
// take time to understand (without recursion)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{0, nil}
	var curr = head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return head.Next
}

func main() {
	var list1 = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	var list2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	var results = mergeTwoLists(list1, list2)

	for results != nil {
		fmt.Println(results.Val)
		results = results.Next
	}
}
