package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (ih IntHeap) Len() int           { return len(ih) }
func (ih IntHeap) Less(i, j int) bool { return ih[i] < ih[j] }
func (ih IntHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }

func (ih *IntHeap) Push(i interface{}) {
	*ih = append(*ih, i.(int))
}

func (ih *IntHeap) Pop() interface{} {
	n := len(*ih)
	x := (*ih)[n-1]
	*ih = (*ih)[0 : n-1]
	return x
}

func main() {
	ih := &IntHeap{4, 1, 2, 5}
	heap.Init(ih)
	heap.Push(ih, 9)
	heap.Push(ih, 6)
	heap.Push(ih, 3)
	for ih.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(ih))
	}
}
