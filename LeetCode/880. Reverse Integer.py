# https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/880/
# Inspired: https://www.geeksforgeeks.org/write-a-program-to-reverse-digits-of-a-number/
class Solution:
    def reverse(self, x: int) -> int:
        if x == 0:
            return 0
        temp = x//abs(x)
        rev = 0
        x = abs(x)
        while x > 0:
            rev = (rev * 10) + (x % 10)
            x = x // 10
        result = rev*temp
        if (result > (2**31 - 1)) or (result < -2**31):
            return 0
        return result
