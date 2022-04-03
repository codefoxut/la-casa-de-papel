package main

import (
	"errors"
	"fmt"
)

type MyStack struct {
	data []int32
}

func (ms *MyStack) Pop() (int32, error) {
	if len(ms.data) > 0 {
		temp := ms.data[len(ms.data)-1]
		ms.data = ms.data[:len(ms.data)-1]
		return temp, nil
	}
	return 0, errors.New("empty stack")
}

func (ms *MyStack) Push(num int32) {
	if ms.data == nil {
		ms.data = make([]int32, 0)
	}
	ms.data = append(ms.data, num)
}

func (ms *MyStack) Top() (int32, error) {
	if len(ms.data) == 0 {
		return 0, errors.New("empty stack")
	}
	return ms.data[len(ms.data)-1], nil
}

func (ms *MyStack) GetLen() int {
	return len(ms.data)
}

func (ms *MyStack) String() {
	fmt.Printf("%v\n", ms.data)
}

func main() {
	mst := new(MyStack)
	v, e := mst.Pop()
	fmt.Println(v, e)
	mst.Push(12)
	mst.Push(13)
	mst.Push(14)
	mst.String()

	v, e = mst.Pop()
	fmt.Println(v, e)

	v, e= mst.Top()
	fmt.Println(v, e)
}
