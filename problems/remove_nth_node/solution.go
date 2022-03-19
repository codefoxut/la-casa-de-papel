package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Two poiters strategy will help in solution.
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

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var nhead, temp *ListNode
	count := 1
	nhead = head
	for count <= n && nhead != nil {
		fmt.Printf("first for %v\t%v\n", count, nhead.GetValue())
		nhead = nhead.GetNext()
		count++
	}

	if nhead == nil {
		return head.GetNext()
	}

	temp = head
	fmt.Printf("before for %v\t%v\n", temp, nhead)
	for nhead != nil {
		if nhead.GetNext() == nil {
			break
		}
		temp = temp.GetNext()
		nhead = nhead.GetNext()
		fmt.Printf("for %v\t%v\n", temp, nhead)
	}
	next := temp.GetNext().GetNext()
	temp.GetNext().SetNext(nil)
	temp.SetNext(next)

	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nhead := NewNode(0)
	nhead.SetNext(head)
	count := 0

	first := nhead
	second := nhead
	for count <= n && nhead != nil {
		// fmt.Printf("for %v\t%v\n", count, first.GetValue())
		first = first.GetNext()
		count++
	}

	for first != nil {
		// fmt.Printf("choose for %v\t%v\n", first.GetValue(), second.GetValue())
		first = first.GetNext()
		second = second.GetNext()

	}
	second.SetNext(second.GetNext().GetNext())

	return nhead.GetNext()
}

func main() {
	listA := []int{5, 7, 9}

	listB := []int{5, 6}

	l1 := CreateList(listA)
	l2 := CreateList(listB)
	l1.Print()
	l2.Print()
	l3 := removeNthFromEnd2(l1, 2)
	l3.Print()

	l4 := removeNthFromEnd(l2, 2)
	l4.Print()
}
