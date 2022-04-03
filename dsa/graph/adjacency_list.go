package main1

import (
	"fmt"
)

type MyUndirectedGraphAL struct {
	data map[int][]int
}

func (mg *MyUndirectedGraphAL) InsertEdge(v1, v2 int) {
	if mg.data == nil {
		mg.data = make(map[int][]int)
	}
	if _, ok := mg.data[v1]; !ok {
		mg.data[v1] = make([]int, 0)
	}
	if _, ok := mg.data[v2]; !ok {
		mg.data[v2] = make([]int, 0)
	}
	mg.data[v1] = append(mg.data[v1], v2)
	mg.data[v2] = append(mg.data[v2], v1)
}

func (mg *MyUndirectedGraphAL) String() {
	for k, v := range mg.data {
		fmt.Printf("%d=>%v\n", k, v)
	}

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

func main() {
	g := new(MyUndirectedGraphAL)
	g.InsertEdge(1, 5)
	g.InsertEdge(5, 2)
	g.InsertEdge(2, 7)
	g.InsertEdge(7, 1)
	g.String()

	d := new(MyDirectedGraphAL)
	d.InsertEdge(1, 5)
	d.InsertEdge(5, 2)
	d.InsertEdge(2, 7)
	d.InsertEdge(7, 1)
	d.String()
}
