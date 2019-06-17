package main

// 実戦で必要になるかわからないがsliceの指し示すarray自体をコピーしたい時

import (
	"fmt"
)

func main() {
	srcSlice := []int{0, 1, 2}
	destSlice := make([]int, len(srcSlice))
	fmt.Println(destSlice)
	copy(destSlice, srcSlice)
	fmt.Println(srcSlice, destSlice)
	srcSlice[0] = 3
	fmt.Println(srcSlice, destSlice)
}
