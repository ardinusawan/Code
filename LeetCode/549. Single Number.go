func singleNumber(nums []int) int {
	// use xor. Same == 0, different == 1
	// 1^1 == 0
	// 1^0 == 1

	v := 0
	for _, n := range nums {
		v ^= n
	}
	return v
}
