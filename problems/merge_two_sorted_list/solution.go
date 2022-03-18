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

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	head := NewNode(0)
	temp := head

	fmt.Printf("%v\t%v\t%v\n", list1, list2, head)
	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list1.GetValue() > list2.GetValue()) {
			temp.SetNext(list2)
			list2 = list2.GetNext()
		} else {
			temp.SetNext(list1)
			list1 = list1.GetNext()
		}

		temp = temp.GetNext()
	}

	return head.GetNext()
}

func main() {
	listA := []int{2}
	listB := []int{1}

	l1 := CreateList(listA)
	l2 := CreateList(listB)
	l1.Print()
	l2.Print()
	fmt.Printf("%v\t%v\n", l1, l2)
	l3 := mergeTwoLists(l1, l2)
	l3.Print()
}
