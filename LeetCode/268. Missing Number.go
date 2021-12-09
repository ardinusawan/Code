// https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {
    // using XOR
    // in binary, XOR will make different value to 1 and same value to 0. 01 ^ 01 = 0, 01 ^ 10 = 11  
    
    A := nums[0]
    for i:=1;i<len(nums);i++ {
        A = nums[i] ^ A
    }
    
    B := 0
    for i:=1;i<=len(nums);i++ {
        B = i ^ B
    }
    
    return A ^ B
}
