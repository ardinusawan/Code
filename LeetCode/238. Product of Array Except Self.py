class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = []
        for i, n in enumerate(nums):
            if i == 0:
                result.append(n)
            else:
                value = result[i-1]*n
                result.append(value)
                
        product = 1
        for i in range(len(nums), 0, -1):
            if i == len(nums):
                result[i-1] = result[i-2]
                product = nums[i-1]
            elif i == 1:
                result[0] = product
            else:
                result[i-1] = result[i-2]*product
                product=product*nums[i-1]
            
        return result
