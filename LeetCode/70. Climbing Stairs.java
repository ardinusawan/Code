// https://leetcode.com/problems/climbing-stairs/

import java.util.HashMap;

class Solution {
    private HashMap<Integer, Integer> memo = new HashMap<>();
    
    public int climbStairs(int n) {
        //pattern is fibonacci
        return calculate(n);
    }
    
    private int calculate(int n) {
        if(n <= 2) {
            return n;
        }
        
        if(memo.containsKey(n)) return memo.get(n);
        
        int result = 0;
        result = calculate(n-1) + calculate(n-2);
        memo.put(n, result);
        return result;
    }
}
