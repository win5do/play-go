class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m = len(word1)
        n = len(word2)

        if m * n == 0:
            return m + n

        D = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m + 1):
            D[i][0] = i
        for j in range(n + 1):
            D[0][j] = j

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                x = D[i - 1][j]
                y = D[i][j - 1]
                z = D[i - 1][j - 1]
                if word1[i - 1] == word2[j - 1]:
                    z -= 1
                D[i][j] = min(x, y, z) + 1

        return D[m][n]


if __name__ == '__main__':
    print(Solution().minDistance('horse', 'ros'))
