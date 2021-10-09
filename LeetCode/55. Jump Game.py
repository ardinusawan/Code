# https://leetcode.com/problems/jump-game/discuss/?currentPage=1&orderBy=hot&query=

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums)==1:
            return True
        left_pointer = nums[0]
        i = 0
        while(left_pointer):
            left_pointer-=1
            if i == len(nums)-1:
                return True
            left_pointer = max(left_pointer, nums[i])
            i+=1
        return False
