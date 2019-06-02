package main

import "fmt"

func binarySearch(nums []int, target int) (int, bool) {
	i := len(nums) / 2
	for {
		if nums[i] == target {
			return i, true
		}

		if i >= len(nums)-1 || i == 0 {
			break
		}

		if nums[i] > target {
			nums = nums[:i]
			i = len(nums) / 2
		} else {
			i = i + (len(nums)-i)/2
		}
	}
	return -1, false
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(s, 4))
	fmt.Println(s)
}
