package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	l, r := mergeSort(nums[:mid]), mergeSort(nums[mid:])
	return merge(l, r)
}

func merge(nums1, nums2 []int) []int {
	i, j := 0, 0
	res := make([]int, 0, len(nums1)+len(nums2))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	} else if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}

func main() {
	nums := []int{4, 5, 1, 8, 3, 2, 9, 6, 7}
	fmt.Println(nums)
	fmt.Println(mergeSort(nums))
}
