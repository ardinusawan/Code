// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
    // Buttom up solution
    // There is a pattern. For every multiple of 2, bit will longer and msp(most significant value, head of bit) is 1
    // current value = 1 + memo[i - exp]
    // int  bit     result
    // 0    = 00    -> 0 
    // 1    = 01    -> 1 -> exp = 1 -> 1 + memo[1 - 1] = 1
    // 2    = 10    -> 1 -> exp = 1 x 2 -> 1 + memo[2 - 2] = 1 
    // 3    = 11    -> 1
    //--------------------------------
    // 4    = 100   -> 1 (exp == 2x2 = 4)
    // In for each i to n, if i is multiple of 2, then exp = prev exp x 2
    
    memo := make([]int, n+1)
    exp := 1
    for i:=1;i<=n;i++ {    
        if i == 2 * exp {
            exp *= 2
        }
        memo[i] = 1 + memo[i-exp]
    }
    return memo
}
