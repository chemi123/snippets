package main

import (
	"fmt"
)

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2
	left := mergeSort(s[:mid])
	right := mergeSort(s[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	l, r := 0, 0
	merge := make([]int, 0, len(left)+len(right))
	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				merge = append(merge, left[l])
				l++
			} else {
				merge = append(merge, right[r])
				r++
			}
			continue
		}

		if l < len(left) {
			merge = append(merge, left[l])
			l++
			continue
		}

		// r < len(right)
		merge = append(merge, right[r])
		r++
	}

	return merge
}

func main() {
	s := []int{4, 2, 6, 9, 1, 3, 7, 8, 5}
	fmt.Println(s)
	fmt.Println(mergeSort(s))
}
