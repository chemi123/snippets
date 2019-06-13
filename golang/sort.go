package main

import (
	"fmt"
	"sort"
)

type Sample struct {
	val   int
	label string
}

func main() {
	s := []Sample{Sample{val: 9, label: "s1"},
		Sample{val: 5, label: "s2"},
		Sample{val: 2, label: "s3"},
		Sample{val: 7, label: "s4"},
		Sample{val: 1, label: "s5"},
		Sample{val: 11, label: "s6"},
		Sample{val: 8, label: "s7"}}
	sort.Slice(s, func(i, j int) bool {
		return s[i].val < s[j].val
	})
	fmt.Println(s)
}
