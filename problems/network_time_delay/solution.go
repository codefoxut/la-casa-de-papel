package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	mapOfCost := make(map[int][][]int)

	for _, val := range times {
		if mapOfCost[val[0]] == nil {
			mapOfCost[val[0]] = make([][]int, 0)
		}
		mapOfCost[val[0]] = append(mapOfCost[val[0]], val[1:])
	}
	h := new(IntHeap)

	heap.Init(h)
	heap.Push(h, []int{0, k})
	// fmt.Printf("minimum: %d\n", (*h)[0])
	visited := make(map[int]bool)
	distance := make([]int, n+1)

	for i := 1; i <= n; i++ {
		distance[i] = int(math.MaxInt8)
	}
	distance[k] = 0

	for h.Len() > 0 {
		node := heap.Pop(h)
		intNode, _ := node.([]int)
		if val := visited[intNode[1]]; val {
			continue
		}
		visited[intNode[1]] = true

		if len(visited) == n {
			return intNode[0]
		}

		// fmt.Printf("%v\n heap %v\ndistance%v\n", visited, h, distance)

		for _, item := range mapOfCost[intNode[1]] {
			if exists := visited[item[0]]; !exists && intNode[0]+item[1] < distance[item[0]] {

				distance[item[0]] = intNode[0] + item[1]
				heap.Push(h, []int{distance[item[0]], item[0]})
			}
		}

	}
	return -1

}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2

	fmt.Println(networkDelayTime(times, n, k))

}
