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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sol := NewNode(0)
	head := sol
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		val := l1.GetValue() + l2.GetValue() + carry
		sol.SetNext(NewNode(val % 10))
		carry = val / 10
		sol = sol.GetNext()
		l1 = l1.GetNext()
		l2 = l2.GetNext()
	}
	return head.GetNext()
}

func addTwoNumbersRaw(l1 *ListNode, l2 *ListNode) *ListNode {
	sol := &ListNode{
		Val: 0,
	}
	head := sol
	var carry, val int
	for l1 != nil || l2 != nil || carry > 0 {
		val = carry

		if l1 != nil {
			val += l1.Val
		}

		if l2 != nil {
			val += l2.Val
		}

		sol.Next = &ListNode{
			Val: val % 10,
		}
		carry = val / 10
		sol = sol.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}
	return head.Next
}

func main() {
	listA := []int{5, 7, 9}

	listB := []int{5, 5}

	l1 := CreateList(listA)
	l2 := CreateList(listB)
	l1.Print()
	l2.Print()
	l3 := addTwoNumbers(l1, l2)
	l3.Print()

	l4 := addTwoNumbersRaw(l1, l3)
	l4.Print()
}
