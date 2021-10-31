# https://leetcode.com/problems/minimum-deletion-cost-to-avoid-repeating-letters

# algo: calculate all cost of repeating letters
# then decreased it with max cost. You will got efficient cost to delete
# use for loop to walk, and use pointer to mark latest repeated char
# jump for loop to the pointer, so we not calculate it more than once

class Solution:
    def minCost(self, s: str, cost: List[int]) -> int:
        result = 0
        i = 0
        while i < (len(s)-1):
            rp = i # pointer use to walk on same char
            total_temp = 0
            max_temp = 0
            while rp < len(s) and s[i] == s[rp]:
                total_temp+=cost[rp] # calculate all cost repeating char
                max_temp = max(max_temp, cost[rp]) # find max temp, use it later to decreasing it with toal temp
                rp+=1
            result+=total_temp-max_temp
            i=rp
        return result
