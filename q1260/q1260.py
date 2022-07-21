from typing import List


class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        for i, row in enumerate(grid):
            for j, v in enumerate(row):
                index1 = (i * n + j + k) % (m * n)
                ans[index1 // n][index1 % n] = v
        return ans


if __name__ == "__main__":
    s = Solution()
    grid = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    print(s.shiftGrid(grid, 1))
