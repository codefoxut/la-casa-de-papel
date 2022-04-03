package main

type Node struct {
	value    int
	previous *Node
	next     *Node
}

func (n *Node) SetNext(node *Node) {
	if n != nil {
		n.next = node
	}
}

func (n *Node) GetNext() *Node {
	if n != nil {
		return n.next
	}
	return nil
}

func (n *Node) SetPrevious(node *Node) {
	if n != nil {
		n.previous = node
	}
}

func (n *Node) GetPrevious() *Node {
	if n != nil {
		return n.previous
	}
	return nil
}

func (n *Node) GetValue() int {
	if n != nil {
		return n.value
	}
	return 0
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}
