package main

func containsDuplicate(nums []int) bool {
	// create hash
	// key is num, value is how much is appears
	// do for loop, for every element increase size of value by one.
	// before increase, do check if value already 1. break : continue

	numH := make(map[int]int, len(nums))

	for _, v := range nums {
		if numH[v] == 1 {
			return true
		}
		numH[v]++
	}
	return false
}
