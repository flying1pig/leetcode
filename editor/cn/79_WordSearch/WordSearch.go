package main

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



 示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCCED"
输出：true


 示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"SEE"
输出：true


 示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCB"
输出：false




 提示：


 m == board.length
 n = board[i].length
 1 <= m, n <= 6
 1 <= word.length <= 15
 board 和 word 仅由大小写英文字母组成




 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

 Related Topics 深度优先搜索 数组 字符串 回溯 矩阵 👍 1968 👎 0

*/

/*
题型: 回溯
题目: 单词搜索
*/

// leetcode submit region begin(Prohibit modification and deletion)
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 匹配失败
			return false
		}
		if k == len(word)-1 { // 匹配成功
			return true
		}
		board[i][j] = 0 // 标记访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y // 相邻格子
			if 0 <= x && x < m && 0 <= y && y < n && dfs(x, y, k+1) {
				return true // 搜到了！
			}
		}
		board[i][j] = word[k] // 恢复现场
		return false          // 没搜到
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true // 搜到了！
			}
		}
	}
	return false // 没搜到
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
