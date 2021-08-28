from typing import List


class Solution:
    # 动态规划
    # 状态转移方程: f(n) = max{ f(n-1) + a[n], a[n] }
    def maxSubArray(self, nums: List[int]) -> int:
        r = nums[0]
        for i in range(1, len(nums)):
            if nums[i] + nums[i - 1] > nums[i]:
                nums[i] += nums[i - 1]
            r = max(r, nums[i])
        return r
