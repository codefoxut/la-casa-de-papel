package main

import (
	"errors"
	"fmt"
)

type MyQueue struct {
	data []int32
}

func (mq *MyQueue) Enqueue(num int32) {
	if mq.data == nil {
		mq.data = make([]int32, 0)
	}
	mq.data = append(mq.data, num)
}

func (mq *MyQueue) Dequeue() (int32, error) {
	if len(mq.data) == 0 {
		return 0, errors.New("empty queue")
	}
	temp := mq.data[0]
	mq.data = mq.data[1:]
	return temp, nil
}

func (mq *MyQueue) Front() (int32, error) {
	if len(mq.data) == 0 {
		return 0, errors.New("empty queue")
	}
	return mq.data[0], nil
}

func (mq *MyQueue) String() {
	fmt.Printf("%v\n", mq.data)
}

func main() {
	mqt := new(MyQueue)
	v, e := mqt.Front()
	fmt.Println(v, e)

	mqt.Enqueue(14)
	mqt.Enqueue(19)
	mqt.Enqueue(1)

	mqt.String()
	v, e = mqt.Dequeue()
	fmt.Println(v, e)
	v, e = mqt.Dequeue()
	fmt.Println(v, e)
}
