import math
import unittest


class Solution:
    # bisection method  二分法
    def mySqrt(self, x: int) -> int:
        left, right = 0, x
        ans = -1

        while left <= right:
            m = (left + right) // 2
            n = m * m

            if n <= x:
                ans = m
                left = m + 1
            else:
                right = m - 1

        return ans


class Solution2:
    # Newton's method, 牛顿法
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0

        C, x0 = float(x), float(x)

        while True:
            xi = 0.5 * (x0 + C / x0)
            if abs(x0 - xi) < 1e-7:
                break
            x0 = xi
        return int(x0)


class Test(unittest.TestCase):
    def test_bisection(self):
        for x in range(0, 1000):
            self.assertEqual(Solution().mySqrt(x), int(math.sqrt(x)))

    def test_newton(self):
        for x in range(0, 1000):
            self.assertEqual(Solution2().mySqrt(x), int(math.sqrt(x)))


if __name__ == '__main__':
    unittest.main()
