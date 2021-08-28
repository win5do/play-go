class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        map = {}
        left = -1
        r = 0

        for i, v in enumerate(s):
            if v in map and map[v] > left:
                left = map[v]
            r = max(r, i - left)
            map[v] = i

        return r


if __name__ == '__main__':
    r = Solution().lengthOfLongestSubstring('abcabcbb')
    print(r)
