package main

import (
	"fmt"
	"math"
)

/**
给你一个下标从 0 开始、大小为 m x n 的矩阵 grid ，矩阵由若干 正 整数组成。

 你可以从矩阵第一列中的 任一 单元格出发，按以下方式遍历 grid ：


 从单元格 (row, col) 可以移动到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1) 三个
单元格中任一满足值 严格 大于当前单元格的单元格。


 返回你在矩阵中能够 移动 的 最大 次数。



 示例 1：
 输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
输出：3
解释：可以从单元格 (0, 0) 开始并且按下面的路径移动：
- (0, 0) -> (0, 1).
- (0, 1) -> (1, 2).
- (1, 2) -> (2, 3).
可以证明这是能够移动的最大次数。

 示例 2：


输入：grid = [[3,2,4],[2,1,9],[1,1,7]]
输出：0
解释：从第一列的任一单元格开始都无法移动。




 提示：


 m == grid.length
 n == grid[i].length
 2 <= m, n <= 1000
 4 <= m * n <= 10⁵
 1 <= grid[i][j] <= 10⁶


 👍 76 👎 0

*/

/*
题型: dp
题目: 矩阵中移动的最大次数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxMoves(grid [][]int) int {
	var f func(int, int, int, int) int
	m := len(grid)
	n := len(grid[0])
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}

	f = func(i int, j int, pre int, depth int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] <= pre {
			return depth
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		a := f(i-1, j+1, grid[i][j], depth)
		b := f(i, j+1, grid[i][j], depth)
		c := f(i+1, j+1, grid[i][j], depth)
		mem[i][j] = max(a, b, c) + 1
		return mem[i][j]
	}

	maxDepth := 0
	for i := 0; i < m; i++ {
		curr := f(i, 0, math.MinInt, 0)
		maxDepth = max(maxDepth, curr)
	}
	return maxDepth - 1
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)为最大移动次数
状态方程:
	a: f(i,j) = f(i-1,j-1)+1, grid[i][j]>grid[i-1][j-1]
	b: f(i,j) = f(i,j-1)+1, grid[i][j]>grid[i][j-1]
	c: f(i,j) = f(i+1,j-1)+1, grid[i][j]>grid[i+1][j-1]
	f(i,j) = max(a,b,c)
边界条件:
	f(-1,j) = f(i,-1) = 0
	f(i,0) = 0
*/

func main() {
	arr := [][]int{
		[]int{3, 2, 4},
		[]int{2, 1, 9},
		[]int{1, 1, 7},
	}
	fmt.Println(maxMoves(arr))
}
