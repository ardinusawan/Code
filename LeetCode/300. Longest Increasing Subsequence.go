// https://leetcode.com/problems/longest-increasing-subsequence/

import (
    "math"
)

func lengthOfLIS(nums []int) int {
    var result int
    if len(nums) == 0 {
        return result
    }
    
    LIS := make([]int, len(nums))
    // longest element for one value is one
    for i, _ := range LIS {
        LIS[i] = 1
    }
    result = LIS[0]
    
    for i := len(nums)-2; i>=0; i-- {
        for j := i+1; j < len(nums); j++ {
            // since last data LIS must be 1, no need to calculate it
            // compare current value with LIS latter value
            // if current value <  latter value, increase LIS
            if nums[i] < nums[j] {
                // LIS[i] = max(1, 1 + LIS[i+1])
                // why compare with 1? because every element has 1 LIS
                LIS[i] = int(math.Max(float64(LIS[i]), float64(1 + LIS[j])))
                result = int(math.Max(float64(LIS[i]), float64(result)))
            }
        }        
    }
    return result
}
