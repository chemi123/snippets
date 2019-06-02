package main

import (
	"fmt"
)

func binarySearch(a []int, target int) int {
	mid := len(a) / 2
	var result int
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > target:
		result = binarySearch(a[:mid], target)
	case a[mid] < target:
		result = binarySearch(a[mid+1:], target)
	default:
		result = mid
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(s, 0))
}
