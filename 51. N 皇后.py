# n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

# 上图为 8 皇后问题的一种解法。

# 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

# 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/n-queens
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


class Solution:
    def solveNQueens(self, n: int):
        board = [-1] * n
        ans = []
        self.findQueens(board, 0, n, ans)
        return ans

    def findQueens(self, board, i, n, ans):
        if n == i:
            tmp = []
            for x in board:
                tmp.append('.'*x + 'Q' + '.'*(n-x-1))
            ans.append(tmp)
            return True
        canPlace = False
        for j in range(n):
            isOK = True
            for x in range(i):
                if board[x] == j or abs(x - i) == abs(board[x] - j):
                    isOK = False
                    break
            if isOK:
                board[i] = j
                canPlace = self.findQueens(board, i + 1, n, ans)
                board[i] = -1
        return canPlace


if __name__ == "__main__":
    solution = Solution()
    print(solution.solveNQueens(4))
