package main

import "strings"

/**
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

 n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。



 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。





 示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。


 示例 2：


输入：n = 1
输出：[["Q"]]




 提示：


 1 <= n <= 9


 Related Topics 数组 回溯 👍 2277 👎 0

*/

/*
题型: 回溯
题目: N 皇后
*/

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) (ans [][]string) {
	queens := make([]int, n) // 皇后放在 (r,queens[r])
	col := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		// 在 (r,c) 放皇后
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] { // 判断能否放皇后
				queens[r] = c                                    // 直接覆盖，无需恢复现场
				col[c], diag1[r+c], diag2[rc] = true, true, true // 皇后占用了 c 列和两条斜线
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
