# https://leetcode.com/problems/valid-parentheses/

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for d in s:
            stack.append(d)
            if len(stack)>=2:
                if stack[-1] == '}' and stack[-2] == '{':
                    stack.pop()
                    stack.pop()
                elif stack[-1] == ')' and stack[-2] == '(':
                    stack.pop()
                    stack.pop()
                elif stack[-1] == ']' and stack[-2] == '[':
                    stack.pop()
                    stack.pop()

        return len(stack) == 0
