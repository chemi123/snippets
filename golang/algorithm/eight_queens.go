package main

import (
	"fmt"
)

func eightQueens() [][]int {
	res := make([][]int, 0)
	for i := 0; i < 8; i++ {
		queens := placeQueen([]int{}, i)
		if len(queens) == 8 {
			res = append(res, queens)
		}
	}
	return res
}

func placeQueen(queens []int, pos int) []int {
	queens = append(queens, pos)
	if len(queens) == 8 {
		return queens
	}
	for i := 0; i < 8; i++ {
		if isValid(queens, i) {
			nextQueens := placeQueen(queens, i)
			if len(nextQueens) == 8 {
				return nextQueens
			}
		}
	}
	return queens
}

func isValid(queens []int, pos int) bool {
	for row, col := range queens {
		if col == pos {
			return false
		}
		diffX := absDiff(col, pos)
		diffY := absDiff(row, len(queens))
		if diffX == diffY {
			return false
		}
	}
	return true
}

func absDiff(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

func main() {
	fmt.Println(eightQueens())
}
