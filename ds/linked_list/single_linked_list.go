package linked_list

import (
	"fmt"
	"strconv"
	"strings"
)

type SingleLinkedList struct {
	head   *Node
	length int
}

func (sll *SingleLinkedList) SetHead(head *Node) {
	if sll != nil {
		sll.head = head
	}
}

func (sll *SingleLinkedList) GetHead() *Node {
	if sll != nil {
		return sll.head
	}
	return nil
}

func (sll *SingleLinkedList) Print() {
	temp := sll.head
	valueList := make([]string, 0, sll.length)

	for temp != nil {
		valueList = append(valueList, strconv.Itoa(temp.value))
		temp = temp.next
	}

	fmt.Println(strings.Join(valueList, " -> "))

}

func (sll *SingleLinkedList) GetPrev(pos int) *Node {
	if sll == nil || pos == 0 {
		return nil
	}
	count := 1
	temp := sll.GetHead()
	for count < pos {
		fmt.Printf(" Nodes %v\t%v\n", count, temp)
		if temp.GetNext() == nil {
			break
		}
		temp = temp.GetNext()
		count++
	}
	return temp
}

func (sll *SingleLinkedList) InsertNode(val, pos int) {
	node := NewNode(val)
	if pos == 0 {
		temp := sll.head
		node.SetNext(temp)
		sll.SetHead(node)
	} else {
		prev := sll.GetPrev(pos)
		fmt.Printf("previous Node %v\t%v\n", pos, prev)
		temp := prev.GetNext()
		node.SetNext(temp)
		prev.SetNext(node)
	}
	sll.length++
}

func (sll *SingleLinkedList) DeleteNode(key int) {
	temp := sll.GetHead()
	if temp == nil {
		return
	}

	if temp.GetValue() == key {
		sll.SetHead(temp.GetNext())
		temp = nil
		sll.length--
		return
	}

	for temp.GetNext() != nil {
		if temp.GetNext().GetValue() == key {
			p := temp.GetNext()
			temp.SetNext(p.GetNext())
			p.SetNext(nil)
			sll.length--
			return
		}
		temp = temp.GetNext()
	}
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func main() {
	s := NewSingleLinkedList()
	s.SetHead(NewNode(5))

	node2 := NewNode(6)
	node3 := NewNode(3)
	node4 := NewNode(4)

	s.head.next = node2
	node2.next = node3
	node3.next = node4

	s.Print()

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
	s.DeleteNode(4)
	s.Print()

	s.DeleteNode(3)
	s.Print()

}
