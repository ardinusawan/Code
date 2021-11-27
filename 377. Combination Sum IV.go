// https://leetcode.com/problems/combination-sum-iv/

func combinationSum4(nums []int, target int) int {
    memo := make(map[int]int)
    memo[0] = 1 // base case
    
    // buttom up
    for i:=1;i<=target;i++ { // solve sub-problem
        for _, n := range nums { // decision tree
            memo[i] += memo[i-n] // adding up prev value to current memo
        }
    }
    
    return memo[target]
}
