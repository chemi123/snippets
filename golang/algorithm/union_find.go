package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
  以下よりunion findを試してみる。acが取れていることを確認している。
  https://atcoder.jp/contests/atc001/tasks/unionfind_a
*/

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) UnionFind {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, -1)
	}
	return UnionFind{parent: nums}
}

func (uf *UnionFind) FindRoot(x int) int {
	if uf.parent[x] < 0 {
		return x
	}
	uf.parent[x] = uf.FindRoot(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.FindRoot(x)
	y = uf.FindRoot(y)
	if x == y {
		return
	}

	if uf.parent[x] > uf.parent[y] {
		x, y = y, x
	}

	uf.parent[x] += uf.parent[y]
	uf.parent[y] = x
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.FindRoot(x) == uf.FindRoot(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.parent[uf.FindRoot(x)]
}

func (uf UnionFind) GroupNum() int {
	var res int
	for i := range uf.parent {
		if uf.parent[i] < 0 {
			res++
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nqStr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	n, _ := strconv.Atoi(nqStr[0])
	q, _ := strconv.Atoi(nqStr[1])
	uf := NewUnionFind(n)
	for i := 0; i < q; i++ {
		scanner.Scan()
		queryStr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		query := queryStr[0]
		v1, _ := strconv.Atoi(queryStr[1])
		v2, _ := strconv.Atoi(queryStr[2])
		if query == "0" {
			uf.Unite(v1, v2)
		} else {
			if uf.IsSame(v1, v2) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
