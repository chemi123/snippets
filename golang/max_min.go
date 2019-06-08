package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		max = int(math.Max(float64(max), float64(num)))
	}
	return max
}

func min(nums ...int) int {
	min := math.MaxInt32
	for _, num := range nums {
		min = int(math.Min(float64(min), float64(num)))
	}
	return min
}

func main() {
	fmt.Println(max(3, 5, 2, 4, 9, 7))
	fmt.Println(min(3, 5, 2, 4, 9, 7))
}
