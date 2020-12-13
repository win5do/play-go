class Solution:
    # 动态规划
    # 状态转移方程: f(n) = f(n-1) + f(n-2)
    def climbStairs(self, n: int) -> int:
        a = 1
        b = 2
        if n == 1:
            return a
        if n == 2:
            return b

        r = -1
        for i in range(3, n + 1):
            r = a + b
            a = b
            b = r
        return r
