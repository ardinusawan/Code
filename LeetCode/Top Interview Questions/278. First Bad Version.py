# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):

class Solution:
    result = 0
    def firstBadVersion(self, n):
        # global result
        # result = n
        """
        :type n: int
        :rtype: int
        """
        l = 1
        h = n
        m = n
        result = n
        
        while(l<=h):
            m = int(l + (h-l)/2)
            print(m)
            if(isBadVersion(m)):
                result = m
                h = m -1
            else:
                l = m +1
        return result
        # self.find(1, n, int((n-1)/2+1))
        # return result
        
    # def find(self, l, h, m):
    #     global result
    #     if isBadVersion(m):
    #         l = l+1
    #         h = m
    #         m = int(l + (h-l)/2)
    #     else:
    #         l = m
    #         h = h
    #         m = int(l + (h-l)/2)
    #     if l == m == h:
    #         result = m
    #         print(m)
    #         return m
    #     else:
    #         self.find(l, h, m)
        
        
        
