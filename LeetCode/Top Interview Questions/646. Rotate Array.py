# https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        k = k%len(nums)
        right = nums[-k:]
        left = nums[:len(nums)-k]
        new = right + left
        # Somehow (new = right + left) not work!
        # So I iterate it 
        for i, n in enumerate(nums):
            nums[i] = new[i]
        
