package main

import (
	"errors"
	"fmt"
)

type MyQueue struct {
	data []int
}

func (mq *MyQueue) Enqueue(num int) {
	if mq.data == nil {
		mq.data = make([]int, 0)
	}
	mq.data = append(mq.data, num)
}

func (mq *MyQueue) Dequeue() (int, error) {
	if len(mq.data) == 0 {
		return 0, errors.New("empty queue")
	}
	temp := mq.data[0]
	mq.data = mq.data[1:]
	return temp, nil
}

func (mq *MyQueue) Front() (int, error) {
	if len(mq.data) == 0 {
		return 0, errors.New("empty queue")
	}
	return mq.data[0], nil
}

func (mq *MyQueue) String() {
	fmt.Printf("%v\n", mq.data)
}

type MyDirectedGraphAL struct {
	data map[int][]int
}

func (mg *MyDirectedGraphAL) InsertEdge(v1, v2 int) {
	if mg.data == nil {
		mg.data = make(map[int][]int)
	}
	if _, ok := mg.data[v1]; !ok {
		mg.data[v1] = make([]int, 0)
	}
	mg.data[v1] = append(mg.data[v1], v2)
}

func (mg *MyDirectedGraphAL) String() {
	for k, v := range mg.data {
		fmt.Printf("%d=>%v\n", k, v)
	}

}

func (mg *MyDirectedGraphAL) BFS(startNode int) {
	mq := new(MyQueue)
	mq.Enqueue(startNode)
	visited := make(map[int]bool)

	for len(mq.data) > 0 {
		temp, err := mq.Dequeue()
		if err == nil {
			// visit the node
			fmt.Printf("node visited %d\n", temp)
			visited[temp] = true

			// add neighbor to queue.
			for _, item := range mg.data[temp] {
				if val := visited[item]; !val {
					mq.Enqueue(item)
				}

			}
		}

	}
}

func main() {

	d := new(MyDirectedGraphAL)
	d.InsertEdge(2, 1)
	d.InsertEdge(2, 5)
	d.InsertEdge(5, 6)
	d.InsertEdge(5, 8)
	d.InsertEdge(6, 9)
	d.String()

	d.BFS(2)
}
