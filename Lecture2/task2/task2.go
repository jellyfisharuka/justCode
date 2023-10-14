package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}

	mergedList := mergeTwoLists(list1, list2)

	fmt.Print("[")
	printList(mergedList)
	fmt.Println("]")
}

func printList(head *ListNode) {
	current := head
	var elements []string
	for current != nil {
		elements = append(elements, fmt.Sprintf("%d", current.Val))
		current = current.Next
	}
	fmt.Print(strings.Join(elements, ", "))
}
