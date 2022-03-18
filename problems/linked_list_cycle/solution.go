package main

import (
	"fmt"
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

// func (n *ListNode) Print() {
// 	var arr []string

// 	temp := n
// 	for temp != nil {
// 		arr = append(arr, strconv.Itoa(temp.GetValue()))
// 		temp = temp.GetNext()
// 	}
// 	fmt.Println(strings.Join(arr, " -> "))

// }

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

func CreateCycleList(arr []int, pos int) *ListNode {
	head := CreateList(arr)
	var posNode, tail, temp *ListNode
	count := 0
	temp = head
	for temp != nil {
		if count == pos {
			posNode = temp
		}
		if temp.GetNext() == nil {
			tail = temp
		}
		temp = temp.GetNext()
		count++
	}
	if posNode !=  nil {
		tail.SetNext(posNode)
	}
	
	return head
}

func hasCycle(head *ListNode) bool {
	hare := head
	tortoise := head

	for hare != nil && tortoise != nil && hare.GetNext() != nil  {
		hare = hare.GetNext().GetNext()
		tortoise = tortoise.GetNext()
		if hare == tortoise {
			return true
		}

	}
	return false
}

func main() {
	// listA := []int{3, 2, 0, -4}
	listB := []int{1}

	l1 := CreateCycleList(listB, -1)
	// l1.Print()

	fmt.Printf("%v\t%v\n", l1, 1)
	l3 := hasCycle(l1)
	fmt.Println("has cycle", l3)
}
