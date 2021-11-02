// https://leetcode.com/problems/container-with-most-water/
// https://leetcode.com/problems/container-with-most-water/discuss/6100/Simple-and-clear-proofexplanation


import (
    "fmt"
    "math"
)

func maxArea(height []int) int {
    left := 0
    right := len(height)-1
    result := 0
    for left < right {
        minOfTwoPointer := int(math.Min(float64(height[left]), float64(height[right])))
        distance := (right - left)
        result_temp := distance * minOfTwoPointer
        result = int(math.Max(float64(result), float64(result_temp)))
        if height[left] < height[right] {
            left = left + 1
        } else {
            right = right - 1 
        }
    }
    return result
}
