package main

import (
	"fmt"
)

type MyUndirectedGraphAM struct {
	size int
	data [][]int
}

func NewUndirectedGraphAM(numberOfNodes int) *MyUndirectedGraphAM {
	n := numberOfNodes + 1
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	m := &MyUndirectedGraphAM{
		size: n,
		data: matrix,
	}
	return m
}

func (mg *MyUndirectedGraphAM) withinBounds(v1, v2 int) bool {
	return v1 >= 0 && v2 >= 0 && v1 <= mg.size && v2 <= mg.size
}

func (mg *MyUndirectedGraphAM) InsertEdge(v1, v2 int) {
	if mg != nil && mg.data != nil && mg.withinBounds(v1, v2) {
		mg.data[v1][v2] = 1
		mg.data[v2][v1] = 1
	}
}

func (mg *MyUndirectedGraphAM) PrintMatrix() {
	for i, v := range mg.data {
		fmt.Printf("%d => %v\n", i, v)
	}
}

func (mg *MyUndirectedGraphAM) PrintGraph() {
	for i, v := range mg.data {
		for j, val := range v {
			if val == 1 {
				fmt.Printf("%d => %d  \n", i, j)
			}
		}
	}
}

type MyDirectedGraphAM struct {
	size int
	data [][]int
}

func NewMyDirectedGraphAM(numberOfNodes int) *MyDirectedGraphAM {
	n := numberOfNodes + 1
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	m := &MyDirectedGraphAM{
		size: n,
		data: matrix,
	}
	return m
}

func (mg *MyDirectedGraphAM) withinBounds(v1, v2 int) bool {
	return v1 >= 0 && v2 >= 0 && v1 <= mg.size && v2 <= mg.size
}

func (mg *MyDirectedGraphAM) InsertEdge(v1, v2 int) {
	if mg != nil && mg.data != nil && mg.withinBounds(v1, v2) {
		mg.data[v1][v2] = 1
	}
}

func (mg *MyDirectedGraphAM) PrintMatrix() {
	for i, v := range mg.data {
		fmt.Printf("%d => %v\n", i, v)
	}
}

func (mg *MyDirectedGraphAM) PrintGraph() {
	for i, v := range mg.data {
		for j, val := range v {
			if val == 1 {
				fmt.Printf("%d => %d  \n", i, j)
			}
		}
	}
}

func main() {
	g := NewUndirectedGraphAM(5)
	g.InsertEdge(1, 2)
	g.InsertEdge(2, 3)
	g.InsertEdge(4, 5)
	g.PrintGraph()
	g.PrintMatrix()

	d := NewMyDirectedGraphAM(5)
	d.InsertEdge(1, 2)
	d.InsertEdge(2, 3)
	d.InsertEdge(4, 5)
	d.PrintGraph()
}
