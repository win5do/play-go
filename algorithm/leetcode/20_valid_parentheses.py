import unittest


class Solution:
    def isValid(self, s: str) -> bool:
        token = {
            ')': '(',
            '}': '{',
            ']': '[',
        }
        stack = []
        for i in range(len(s)):
            char = s[i]
            if char in token:
                if len(stack) == 0:
                    return False
                top = stack.pop()
                if token[char] != top:
                    return False
            else:
                stack.append(char)

        if len(stack) > 0:
            return False

        return True


class Test(unittest.TestCase):
    def test_isValid(self):
        for i in [
            (True, '([])'),
            (False, '([]'),
            (False, '([)]'),
            (False, '([{'),
        ]:
            self.assertEqual(i[0], Solution().isValid(i[1]))


if __name__ == '__main__':
    unittest.main()
