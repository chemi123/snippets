package main

import (
	"fmt"
	"sort"
)

func searchInts(s []int, num int) int {
	return sort.Search(len(s), func(i int) bool { return s[i] <= num })
}

func main() {
	s := []int{14, 13, 9, 8, 6, 3, 1}
	fmt.Println(sort.SearchInts(s, 3)) // 0
	fmt.Println(searchInts(s, 8))      // 3
}
