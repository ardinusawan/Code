# https://leetcode.com/problems/maximum-product-subarray/
# https://leetcode.com/problems/maximum-product-subarray/discuss/1520026/How-to-come-up-with-a-DP-solution

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        # if there is 0 in the list, probably this is the max
        result = max(nums)
        
        # local min max
        max_sofar = 1
        min_sofar = 1
        
        for n in nums:
            # there is possibility local max is exisiting value, or prev minus local min * -value
            tmax = max(max(max_sofar * n, min_sofar * n), n)
            
            # there is possibility local min is exisiting value, or prev local max * value
            tmin = min(min(max_sofar * n, min_sofar * n), n)
            
            max_sofar = tmax
            min_sofar = tmin
            
            result = max(result, max_sofar)
        return result
