// https://leetcode.com/problems/house-robber/
// https://youtu.be/73r3KWiEvyk

import (
    "math"
)


func rob(nums []int) int {
    var rob1, rob2 int
    for _, n := range nums {
        temp := int(math.Max(float64(n + rob1), float64(rob2)))
        rob1 = rob2
        rob2 = temp
    }
    return rob2
}
