package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Num   int
	Count int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Count > pq[j].Count }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(i interface{}) {
	*pq = append(*pq, i.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return x
}

func main() {
	pq := &PriorityQueue{
		Item{Num: 1, Count: 2},
		Item{Num: 2, Count: 3},
		Item{Num: 3, Count: 1},
	}
	heap.Init(pq)
	heap.Push(pq, Item{Num: 4, Count: 5})
	heap.Push(pq, Item{Num: 3, Count: 2})
	for pq.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(pq))
	}
}
