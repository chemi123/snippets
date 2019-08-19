package main

import (
	"fmt"
	"sort"
)

func SortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

func main() {
	fmt.Println(SortString("Hello, playground"))
}
