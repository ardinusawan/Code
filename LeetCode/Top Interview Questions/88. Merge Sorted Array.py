# https://leetcode.com/problems/merge-sorted-array/

class Solution:
    def merge(self, nums1: 'List[int]', m: 'int', nums2: 'List[int]', n: 'int') -> 'None':
        """
        Do not return anything, modify nums1 in-place instead.
        """
        zero = 0
        len1 = len(nums1)
        while len(nums1) and nums1[-1] == 0:
            zero += 1
            nums1.pop()

        while(len(nums2)):
            if nums1 and nums1[-1] < nums2[0]:
                nums1.append(nums2.pop(0))
            else:
                temp = []
                while(len(nums1) and nums1[-1]>nums2[0]):
                    temp.append(nums1.pop())
                nums1.append(nums2.pop(0))
                while(len(temp)):
                    nums1.append(temp.pop())
        while len1 != len(nums1):
            nums1.append(0)

