# https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/?currentPage=1&orderBy=hot&query=

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # technic: sliding windows
        uniqString = set()
        result = 0
        l = 0
        for r in range(len(s)):
            while s[r] in uniqString:
                uniqString.remove(s[l])
                l += 1
            else:
                uniqString.add(s[r])
            result = max(result, r-l + 1)
        return result
