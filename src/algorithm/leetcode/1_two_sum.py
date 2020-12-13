from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dict = {}
        for k, v in enumerate(nums):
            if (target - v) in dict:
                return [k, dict[target - v]]
            dict[v] = k
        return [-1, -1]


if __name__ == '__main__':
    r = Solution().twoSum([2, 7, 11, 15], 9)
    print(r)
