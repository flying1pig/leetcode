package main

/**
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。

 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。

 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。



 示例 1：


输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。


 示例 2：


输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 500
 grid[i][j] 的值为 0 或 1


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 293 👎 0

*/

/*
题型: 网格dfs
题目: 飞地的数量
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numEnclaves(grid [][]int) int {
	var dx = []int{0, 1, -1, 0}
	var dy = []int{1, 0, 0, -1}
	ans := 0
	f := 0
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return
		}
		if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1 {
			ans = 0
			f = 1
		}
		if f == 0 {
			ans++
		}
		grid[i][j] = 0
		for k := 0; k < 4; k++ {
			x, y := dx[k]+i, dy[k]+j
			dfs(x, y)
		}
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				f = 0
				sum += ans
				ans = 0
			}
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
