// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/


// create 3 pointer, low mid heigh
// low is 0, high is length of array - 1, mid is middle
// split array into 2 based on value in mid
// if value in mid less than value in high, lowest value in low until mid pointer
// else lowest value is in mid + 1 until high
// so if lowest in left of middle, set high = mid
// else low = mid + 1
// repeat until low >= high

func findMin(nums []int) int {
    low := 0
    high := len(nums)-1
    for low < high {
        // find mid
        mid := low + (high - low)/2
        // if value mid < value high, lowest value must be from low to mid
        if nums[mid] < nums[high] {
            high = mid
        } else { // it must be from mid + 1 to high
            low = mid + 1
        }
    }
    return nums[low]
}
