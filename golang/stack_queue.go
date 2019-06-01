package main

import (
	"fmt"
	"math"
)

type Stack []int
type Queue []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() {
	if len(*s) == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Peek() int {
	if len(*s) == 0 {
		return math.MinInt32
	}
	return (*s)[len(*s)-1]
}

func (q *Queue) Push(num int) {
	*q = append(*q, num)
}

func (q *Queue) Pop() {
	if len(*q) == 0 {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) Peek() int {
	if len(*q) == 0 {
		return math.MinInt32
	}
	return (*q)[0]
}

func main() {
	fmt.Println("===Stack===")
	s := make(Stack, 0, 10)
	s.Push(1)
	s.Push(2)
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)

	fmt.Println("\n===Queue===")
	q := make(Queue, 0, 10)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q)
}
