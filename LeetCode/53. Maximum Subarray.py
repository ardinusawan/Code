# https://leetcode.com/problems/maximum-subarray/discuss/?currentPage=1&orderBy=hot&query=

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        result = nums[0]
        temp_result = 0
        for n in nums:
            if temp_result < 0:
                temp_result = 0
            temp_result+=n
            result = max(result, temp_result)
        return result
