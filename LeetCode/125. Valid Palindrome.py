# https://leetcode.com/problems/valid-palindrome/discuss/?currentPage=1&orderBy=hot&query=

class Solution:
    def isPalindrome(self, s: str) -> bool:
        new_char = []
        for c in s:
            if self.is_char(c):
                new_char.append(c.lower())
                
        for i in range(len(new_char)):
            if new_char[i] != new_char[-i-1]:
                return False
        return True 
    def is_char(self, c):
        return c.isalnum()
