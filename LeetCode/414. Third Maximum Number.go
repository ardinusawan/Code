// https://leetcode.com/problems/third-maximum-number/

func thirdMax(nums []int) int {    
    biggest1 := nums[0]
    min := nums[0]

    // find min and max
    for _, n := range nums {
        if biggest1 < n {
            biggest1 = n
        }
        if min > n {
            min = n
        }
    }
    
    // just return biggest if input less than 3 item
    if len(nums) < 3 {
        return biggest1
    }
    
    biggest2 := min
    biggest3 := min

    // biggest2
    for _, n := range nums {
        if n < biggest1 && n >= biggest2 {
            biggest2 = n
        }
    }
    
    // biggest3
    for _, n := range nums {
        if n < biggest2 && n >= biggest3 {
            biggest3 = n
        }
    }
    
    // If the third maximum does not exist, return the maximum number.
    // i.e not unique
    if (biggest1 == biggest2 || biggest2 == biggest3 || biggest1 == biggest3) {
        return biggest1
    } else {
        return biggest3 
    }
}
