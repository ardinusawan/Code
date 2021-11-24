// https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
    dp := make(map[int]bool, len(s)+1)
    // best case
    dp[len(s)] = true
    
    // calculate from backward
    for i:=len(s)-1;i>=0;i--{
        // check foreach wordDict
        for _, w := range wordDict {
            // if word is exist in s
            if (i+len(w)) <= len(s) && s[i:i+len(w)] == w {
                // set current value with prev value in same length word
                // if from tail to top value is true, it correct
                dp[i] = dp[i+(len(w))]
            }
            // if we found match word in s, break it so we continue to check others dict
            if dp[i] {
                break
            }
        }
    }
    
    return dp[0]
}
