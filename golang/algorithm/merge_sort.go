package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	l, r := mergeSort(nums[:mid]), mergeSort(nums[mid:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	res := make([]int, 0, len(l) + len(r))
	var left, right int
	for left < len(l) && right < len(r) {
		if l[left] < r[right] {
			res = append(res, l[left])
			left++
		} else {
			res = append(res, r[right])
			right++
		}
	}
	if left < len(l) {
		res = append(res, l[left:]...)
		return res
	}
	res = append(res, r[right:]...)
	return res
}

func main() {
	nums := make([]int, 30)
	for i := range nums {
		nums[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	fmt.Println(nums)
	fmt.Println(mergeSort(nums))
}