# https://leetcode.com/problems/valid-anagram/

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        sMap, tMap = {}, {}
        for d in s:
            sMap[d] = 1 + sMap.get(d, 0)
        for d in t:
            tMap[d] = 1 + tMap.get(d, 0)
            
        return sMap == tMap
