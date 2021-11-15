// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
    // edge case
    if len(s) < len(t) {
        return ""
    }
    
    // mapping t to key-value
    countT := make(map[string]int)
    
    // mapping s to key-value
    window := make(map[string]int)
    
    for _, v := range t {
        countT[string(v)] += 1 // fill target map
        window[string(v)] = 0 // init
    }
    
    // It used to optimize checking is all key exist and count match
    have, need := 0, len(countT)
    
    // 10001 is enough for upper limit
    result, resLen := []int{0, 0}, 100001
    l := 0
    for r, v := range s {
        window[string(v)] += 1 // fill windows value
        
        // if s exist in t and value is same, it match
        _, exist := countT[string(v)]
        if exist && window[string(v)] == countT[string(v)] {
            have+=1
        }
        
        for have == need {
            // if distance is less than before, set as new result
            if (r - l + 1) < resLen {
                result = []int{l, r+1}
                resLen = (r - l + 1)
            }
            
            window[string(s[l])] -= 1 // remove left pointer value because we will move the pointer
            _, exist := countT[string(s[l])]
            if exist && window[string(s[l])] < countT[string(s[l])] {
                // because we remove left pointer and if value is in target, decrease the have
                have-=1 
            }
            l+=1 // always move left pointer
        }
    }
    return s[result[0]:result[1]]
}
