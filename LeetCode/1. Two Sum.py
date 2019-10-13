# https://leetcode.com/problems/two-sum/
# https://leetcode.com/submissions/detail/259602094/

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        temp = {}
        for v in nums:
            key = target - v
            value = v
            temp[key] = value
        result = []
        for i,v in enumerate(nums):
            if v == target/2 and nums.count(v) == 1:
                pass
            else:
                if v in temp.keys():
                    result.append(i)
        
        return result
