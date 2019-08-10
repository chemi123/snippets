package main

import (
	"fmt"
	"sort"
)

type SortRunes []rune

func (r SortRunes) Len() int           { return len(r) }
func (r SortRunes) Less(i, j int) bool { return r[i] < r[j] }
func (r SortRunes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(SortRunes(r))
	return string(r)
}

func main() {
	s := SortString("string")
	fmt.Println(s)
}
