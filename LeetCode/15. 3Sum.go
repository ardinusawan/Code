// https://leetcode.com/problems/3sum/

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    // sort value so we can create left and right pointer
    sort.Ints(nums)
    var result [][]int 
    for i, n := range nums {
        // if current value == prev value, skip cz we already calculate it
        if i > 0 && n == nums[i-1] {
            continue
        }
        
        // create new pointer left and right, the position if array after n
        l, r := i+1, len(nums)-1
        for l < r {
            three_sum := n + nums[l] + nums[r]
            // if value of three sum more than 0 we need to decreased by moving right pointer to left and vice versa
            if three_sum > 0 {
                r = r - 1
            } else if three_sum < 0 {
                l = l + 1
            } else { // if it 0, it means we got the result
                result = append(result, []int{n, nums[l], nums[r]})
                l = l + 1
                // make sure left pointer is different than before
                for len(nums) > l && nums[l] == nums[l-1] {
                   l = l + 1 
                }
            }
        }
    }
    return result
}
