package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) SetNext(node *ListNode) {
	if n != nil {
		n.Next = node
	}
}

func (n *ListNode) GetNext() *ListNode {
	if n != nil {
		return n.Next
	}
	return nil
}

// func (n *Node) SetPrevious(node *Node) {
// 	if n != nil {
// 		n.previous = node
// 	}
// }

// func (n *Node) GetPrevious() *Node {
// 	if n != nil {
// 		return n.previous
// 	}
// 	return nil
// }

func (n *ListNode) GetValue() int {
	if n != nil {
		return n.Val
	}
	return 0
}

func (n *ListNode) Print() {
	var arr []string

	temp := n
	for temp != nil {
		arr = append(arr, strconv.Itoa(temp.GetValue()))
		temp = temp.GetNext()
	}
	fmt.Println(strings.Join(arr, " -> "))

}

func NewNode(value int) *ListNode {
	return &ListNode{
		Val: value,
	}
}

func CreateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := NewNode(arr[0])
	temp := head
	for i := 1; i < len(arr); i++ {
		temp.SetNext(NewNode(arr[i]))
		temp = temp.GetNext()
	}
	return head

}

func reverseList1(head *ListNode) *ListNode {
	var temp, curr, headNew *ListNode

	temp = head
	headNew = temp
	for temp.GetNext() != nil {
		curr = temp.GetNext()
		temp.SetNext(curr.GetNext())
		curr.SetNext(headNew)
		headNew = curr
		// fmt.Println(temp, curr, headNew)

	}

	return headNew
}

func reverseList(head *ListNode) *ListNode {
	var curr, headNew *ListNode

	headNew = nil
	for head != nil {
		curr = head.GetNext()
		head.SetNext(headNew)
		headNew = head
		head = curr
		// fmt.Println(temp, curr, headNew)

	}

	return headNew
}

func main() {
	listA := []int{1, 3, 5, 4, 2}

	l1 := CreateList(listA)
	l1.Print()
	l2 := reverseList(l1)
	l2.Print()
	l3 := reverseList1(l1)
	l3.Print()
}
