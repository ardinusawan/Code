# Problem: https://leetcode.com/problems/two-sum/
# Detail: https://leetcode.com/submissions/detail/554808992/

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        r = []
        for i, n in enumerate(nums):
            if target-n in m:
                r = [m[target-n], i]
            m[n] = i
        return r
        
