package main

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

 此外，你可以假设该网格的四条边均被水包围。



 示例 1：


输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1


 示例 2：


输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 300
 grid[i][j] 的值为 '0' 或 '1'


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2728 👎 0

*/

/*
题型: 网格dfs
题目: 岛屿数量
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		// 出界，或者不是 '1'，就不再往下递归
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2' // 插旗！避免来回横跳无限递归
		dfs(i, j-1)      // 往左走
		dfs(i, j+1)      // 往右走
		dfs(i-1, j)      // 往上走
		dfs(i+1, j)      // 往下走
	}

	for i, row := range grid {
		for j, c := range row {
			if c == '1' { // 找到了一个新的岛
				dfs(i, j) // 把这个岛插满旗子，这样后面遍历到的 '1' 一定是新的岛
				ans++
			}
		}
	}
	return
}

//时间复杂度：O(mn)，其中 m 和 n 分别为 grid 的行数和列数
//空间复杂度：O(mn)

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
