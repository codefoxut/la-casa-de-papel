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

func oddEvenList(head *ListNode) *ListNode {
	evenHead := head.GetNext()
	oddTail := head
	counter := 1
	temp := head
	curr := head
	for curr != nil {
		counter++
		curr = temp.GetNext()
		next := temp.GetNext().GetNext()
		temp.SetNext(next)
		if counter%2 == 0 {
			oddTail = temp
		}
		temp = curr
	}
	oddTail.SetNext(evenHead)

	return head
}

func oddEvenListOddJump(head *ListNode) *ListNode {
	evenHead := head.GetNext()
	temp := head
	curr := head.GetNext().GetNext()

	for curr != nil {
		curr = temp.GetNext().GetNext()
		next := temp.GetNext()
		if curr == nil {
			temp.SetNext(evenHead)
			break
		}
		temp.SetNext(curr)
		next.SetNext(curr.GetNext())
		if curr.GetNext() == nil {
			curr.SetNext(evenHead)
			break
		}

		temp = curr
	}

	return head
}

func oddEvenListJump(head *ListNode) *ListNode {
	if head == nil || head.GetNext() == nil {
		return head
	}
	temp := head
	tempEven := head.GetNext()

	for temp != nil && tempEven != nil {
		// nextOdd := tempEven.GetNext()
		evenhead := temp.GetNext()
		if tempEven.GetNext() != nil {
			temp.SetNext(tempEven.GetNext())
			tempEven.SetNext(tempEven.GetNext().GetNext())
			temp.GetNext().SetNext(evenhead)
		}

		temp = temp.GetNext()
		tempEven = tempEven.GetNext()
	}

	return head
}

func main() {
	listA := []int{5, 7, 9, 2, 4, 6}

	listB := []int{1, 2, 3, 4, 5, 6, 7}

	l1 := CreateList(listA)
	l2 := CreateList(listB)
	l1.Print()
	l2.Print()
	l3 := oddEvenList(l1)
	l3.Print()

	l4 := oddEvenListOddJump(l2)
	l4.Print()

	l5 := oddEvenListJump(l3)
	l5.Print()

	l6 := oddEvenListJump(l4)
	l6.Print()
}
