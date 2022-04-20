package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	// make hash of nums1
	// key = num
	// value = total value showed up

	// make for loop of nums2
	// for every n, if exist in hash
	// 	 append to result
	// 	 decrase value of hash

	numOneH := make(map[int]int)
	for _, n := range nums1 {
		numOneH[n]++
	}

	for _, n := range nums2 {
		if numOneH[n] > 0 {
			result = append(result, n)
			numOneH[n]--
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersect(nums1, nums2)
	fmt.Println(result)

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	result1 := intersect(nums3, nums4)
	fmt.Println(result1)
}
