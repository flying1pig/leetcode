package main

/**
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。





 示例 1：


输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。




 示例 2：


输入：n = 1
输出：1




 提示：


 1 <= n <= 9


 Related Topics 回溯 👍 565 👎 0

*/

/*
题型: 回溯
题目: N 皇后 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) (ans int) {
	col := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			ans++ // 找到一个合法方案
			return
		}
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] {
				col[c], diag1[r+c], diag2[rc] = true, true, true
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
