package main

/**
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。

 请返回 封闭岛屿 的数目。



 示例 1：




输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,
1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。

 示例 2：




输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1


 示例 3：


输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2




 提示：


 1 <= grid.length, grid[0].length <= 100
 0 <= grid[i][j] <=1


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 317 👎 0

*/

/*
题型: 网格dfs
题目: 统计封闭岛屿的数目
*/

// leetcode submit region begin(Prohibit modification and deletion)
func closedIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
			return
		}
		grid[x][y] = 1 // 标记 (x,y) 被访问，避免重复访问
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for i := 0; i < m; i++ {
		// 如果是第一行和最后一行，访问所有格子
		// 如果不是，只访问第一列和最后一列的格子
		step := 1
		if 0 < i && i < m-1 {
			step = n - 1
		}
		for j := 0; j < n; j += step {
			dfs(i, j)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 { // 从没有访问过的 0 出发
				ans++ // 一定是封闭岛屿
				dfs(i, j)
			}
		}
	}
	return
}

//时间复杂度: o(mn)
//空间复杂度: o(mn)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
