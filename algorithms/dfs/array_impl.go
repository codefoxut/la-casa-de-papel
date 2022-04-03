package main

import (
	"errors"
	"fmt"
)

type MyStack struct {
	data []int
}

func (ms *MyStack) Pop() (int, error) {
	if len(ms.data) > 0 {
		temp := ms.data[len(ms.data)-1]
		ms.data = ms.data[:len(ms.data)-1]
		return temp, nil
	}
	return 0, errors.New("empty stack")
}

func (ms *MyStack) Push(num int) {
	if ms.data == nil {
		ms.data = make([]int, 0)
	}
	ms.data = append(ms.data, num)
}

func (ms *MyStack) Top() (int, error) {
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

func (mg *MyDirectedGraphAL) DFS(startNode int) {
	stack := new(MyStack)
	stack.Push(startNode)
	visited := make(map[int]bool)

	for len(stack.data) > 0 {
		temp, err := stack.Pop()
		if err == nil {
			fmt.Println("node visited", temp)
			visited[temp] = true
			// neighbors
			for _, item := range mg.data[temp] {
				if value := visited[item]; !value {
					stack.Push(item)
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

	d.DFS(2)
}
