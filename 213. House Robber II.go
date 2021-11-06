//  https://leetcode.com/problems/house-robber-ii/

import (
    "math"
)

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    // compare max from first data until n-1 data
    // with second data to until n data
    return int(math.Max(float64(helper(nums[:len(nums)-1])), float64(helper(nums[1:]))))
}

// this is same logic with house robber 1
func helper(nums []int) int {
    rob1, rob2 := 0, 0 
    for _, n := range nums {
        temp := int(math.Max(float64(n + rob1), float64(rob2)))
        rob1 = rob2
        rob2 = temp
    }
    return rob2
}
