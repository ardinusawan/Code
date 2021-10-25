# https://leetcode.com/problems/longest-consecutive-sequence/
# idea: https://medium.com/nerd-for-tech/longest-consecutive-sequence-leetcode-128-medium-1adea28af7a5

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        result = 0
        # set used to check if element exist it cheap and fast way
        nums_set = set(nums)
        # make list not contains duplicate element
        nums = list(nums_set)
        for n in nums:
            # skip if there is lower point than current value
            # we want to get smallest possible value and start counting up from it
            if n-1 in nums_set:
                continue
            # every value is have 1 length right?
            result_temp = 1
            tn = n
            # if value increment is always exist, increment result
            while((tn+1) in nums_set):
                tn=tn+1
                result_temp+=1
            result = max(result, result_temp)
        return result
