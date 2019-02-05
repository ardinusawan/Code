# https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/

class Solution:
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        init = 99999
        for i, v in reversed(list(enumerate(nums))):
            if v == init:
                deleted = nums.pop(i)
            elif v < init:
                init = v

