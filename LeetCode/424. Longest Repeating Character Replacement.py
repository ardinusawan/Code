# https://leetcode.com/problems/longest-repeating-character-replacement/

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        count = {}
        l_ptr = 0
        result = 0
        for r_ptr in range(len(s)):
            count[s[r_ptr]] = 1 + count.get(s[r_ptr], 0)
            if((r_ptr - l_ptr + 1) - max(count.values())) > k: # if total char need to be changed more than k
                count[s[l_ptr]] -= 1 # decrase max value
                l_ptr+=1 # slide left ptr to right
            result = max(result, r_ptr - l_ptr + 1)
        return result
