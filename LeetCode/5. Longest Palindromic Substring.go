// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
    // create 2 pointer, left and right
    // calculate it twice
    // 1. find polindrom of odd length
    // 2. find polindrom of even length
    
    lOdd, rOdd := 0, 0
    for i, _ := range s {
        l, r := i, i
        for l >= 0 && r < len(s) {
            if s[l] == s[r] {
                if r - l > rOdd - lOdd {
                   lOdd, rOdd = l, r 
                }
            } else {
                break
            }
            l-=1
            r+=1
        }
    }
    
    lEven, rEven := 0, 0
    for i, _ := range s {
        l, r := i, i+1
        for l >= 0 && r < len(s) {
            if s[l] == s[r] {
                if r - l > rEven - lEven {
                   lEven, rEven = l, r 
                }
            } else {
                break
            }
            l-=1
            r+=1
        }
    }
    
    var l, r int
    fmt.Println(lOdd, rOdd)
    fmt.Println(lEven, rEven)
    if rOdd - lOdd > rEven - lEven {
        l, r = lOdd, rOdd
    } else {
        l, r = lEven, rEven
    }
    return s[l:r+1] // +1 because char itself is polindrom
}
