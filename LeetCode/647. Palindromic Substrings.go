// https://leetcode.com/problems/palindromic-substrings/

func countSubstrings(s string) int {
    // calculate it twice
    // 1. calculate using left and right pointer in same start location [0] index
        // it only calculate odds polindrom value
        // n^2
        // [a, a, a] -> [a, a, a, a, a] -> ...
    // 2. calculate using left and right pointer in different start
        // left = [0], right = [1]
        // n^2
        // [a, a] -> [a, a, a, a] -> ...
    // n^2 + n^2 = n^2
   
    if len(s) == 0 {
        return 0
    } else if len(s) == 1 {
        return 1
    }
    
    var result int
    // 1
    for i, _ := range s {
        l, r := i, i
        for l >= 0 && r <= len(s)-1 {
            if s[l] == s[r] {
                result+=1
                l-=1
                r+=1
            } else {
                break
            }
        }
    }
    
    // 2
    for i, _ := range s {
        l, r := i, i+1
        for l >= 0 && r <= len(s)-1 {
            if s[l] == s[r] {
                result+=1
                l-=1
                r+=1
            } else {
                break
            }
        }
    }
     return result
}
