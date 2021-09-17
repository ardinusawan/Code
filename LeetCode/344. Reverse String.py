# https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/

class Solution:
    def reverseString(self, s: 'List[str]') -> 'None':
        """
        Do not return anything, modify s in-place instead.
        """
        for i in range(int(len(s)/2)):
            s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]

