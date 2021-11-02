// https://leetcode.com/problems/search-in-rotated-sorted-array/
// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/741586/Visual-explanation


func search(nums []int, target int) int {
    low := 0
    high := len(nums)-1
    for low <= high {
        if nums[low] == target { return low }
        if nums[high] == target { return high }
        mid := low + (high - low)/2
        if nums[mid] == target { return mid }
        
        if nums[mid] > nums[high] {
            if target > nums[mid] || target < nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        } else {
            if target > nums[mid] && target < nums[high] {
                low = mid + 1
            } else {
                high = mid -1
            }
        }
    }
    return -1
}

