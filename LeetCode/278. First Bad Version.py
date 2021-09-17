# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        left = 1
        right = n
        mid = 1
        result = n
        while(left<right):
            mid = int(left + (right-left)/2)
            if isBadVersion(mid):
                result = mid
                right = mid
            else:
                left = mid +1
        return result
            
        
