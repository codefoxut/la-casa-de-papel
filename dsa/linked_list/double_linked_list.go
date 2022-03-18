package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DoubleLinkedList struct {
	head   *Node
	length int
}

func (sll *DoubleLinkedList) SetHead(head *Node) {
	if sll != nil {
		sll.head = head
	}
}

func (sll *DoubleLinkedList) GetHead() *Node {
	if sll != nil {
		return sll.head
	}
	return nil
}

func (sll *DoubleLinkedList) Print() {
	temp := sll.head
	valueList := make([]string, 0, sll.length)

	for temp != nil {
		valueList = append(valueList, strconv.Itoa(temp.value))
		temp = temp.next
	}

	fmt.Println(strings.Join(valueList, " -> "))

}

func (sll *DoubleLinkedList) InsertNode(val, pos int) {
	node := NewNode(val)
	temp := sll.head
	if pos == 0 {
		node.SetNext(temp)
		sll.SetHead(node)
	} else {
		count := 1
		for count < pos && temp.GetNext() != nil {
			temp = temp.GetNext()
			count++
		}

		fmt.Printf("pos Node %v\t%v\n", pos, temp)
		node.SetNext(temp.GetNext())
		temp.GetNext().SetPrevious(node)
		temp.SetNext(node)
		node.SetPrevious(temp)
	}
}

func (sll *DoubleLinkedList) DeleteNode(key int) {
	temp := sll.GetHead()
	if temp == nil {
		return
	}

	if temp.GetValue() == key {
		sll.SetHead(temp.GetNext())
		temp = nil
		return
	}

	for temp.GetNext() != nil {
		if temp.GetNext().GetValue() == key {
			p := temp.GetNext()
			temp.SetNext(p.GetNext())
			p.GetNext().SetPrevious(temp)

			p.SetPrevious(nil)
			p.SetNext(nil)
			return
		}
		temp = temp.GetNext()
	}
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func main() {
	s := NewDoubleLinkedList()
	s.SetHead(NewNode(5))

	s.InsertNode(10, 2)
	s.Print()
	s.InsertNode(11, 1)
	s.Print()
	s.InsertNode(1, 0)
	s.Print()
	s.InsertNode(19, 10)
	s.Print()
	s.InsertNode(18, 10)
	s.Print()
	s.InsertNode(4, 3)
	s.Print()
	s.InsertNode(3, 2)

	s.Print()
	s.DeleteNode(4)
	s.Print()

	s.DeleteNode(3)
	s.Print()

}
