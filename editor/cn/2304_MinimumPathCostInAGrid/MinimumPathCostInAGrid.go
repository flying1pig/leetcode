package main

import (
	"math"
	"slices"
)

/**
给你一个下标从 0 开始的整数矩阵 grid ，矩阵大小为 m x n ，由从 0 到 m * n - 1 的不同整数组成。你可以在此矩阵中，从一个单元格移动到
 下一行 的任何其他单元格。如果你位于单元格 (x, y) ，且满足 x < m - 1 ，你可以移动到 (x + 1, 0), (x + 1, 1), ...
, (x + 1, n - 1) 中的任何一个单元格。注意： 在最后一行中的单元格不能触发移动。

 每次可能的移动都需要付出对应的代价，代价用一个下标从 0 开始的二维数组 moveCost 表示，该数组大小为 (m * n) x n ，其中
moveCost[i][j] 是从值为 i 的单元格移动到下一行第 j 列单元格的代价。从 grid 最后一行的单元格移动的代价可以忽略。

 grid 一条路径的代价是：所有路径经过的单元格的 值之和 加上 所有移动的 代价之和 。从 第一行 任意单元格出发，返回到达 最后一行 任意单元格的最小路径
代价。



 示例 1：




输入：grid = [[5,3],[4,0],[2,1]], moveCost = [[9,8],[1,5],[10,12],[18,6],[2,4],[14,
3]]
输出：17
解释：最小代价的路径是 5 -> 0 -> 1 。
- 路径途经单元格值之和 5 + 0 + 1 = 6 。
- 从 5 移动到 0 的代价为 3 。
- 从 0 移动到 1 的代价为 8 。
路径总代价为 6 + 3 + 8 = 17 。


 示例 2：


输入：grid = [[5,1,2],[4,0,3]], moveCost = [[12,10,15],[20,23,8],[21,7,1],[8,1,13],
[9,10,25],[5,3,2]]
输出：6
解释：
最小代价的路径是 2 -> 3 。
- 路径途经单元格值之和 2 + 3 = 5 。
- 从 2 移动到 3 的代价为 1 。
路径总代价为 5 + 1 = 6 。



 提示：


 m == grid.length
 n == grid[i].length
 2 <= m, n <= 50
 grid 由从 0 到 m * n - 1 的不同整数组成
 moveCost.length == m * n
 moveCost[i].length == n
 1 <= moveCost[i][j] <= 100


 Related Topics 数组 动态规划 矩阵 👍 91 👎 0

*/

/*
题型: dp
题目: 网格中的最小路径代价
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[m-1] = grid[m-1]
	for i := m - 2; i >= 0; i-- {
		for j, g := range grid[i] {
			f[i][j] = math.MaxInt
			for k, c := range moveCost[g] { // 移动到下一行的第 k 列
				f[i][j] = min(f[i][j], f[i+1][k]+c)
			}
			f[i][j] += g
		}
	}
	return slices.Min(f[0])
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
						n−1
	dfs(i,j)=grid[i][j]+min dfs(i+1,k)+moveCost[grid[i][j]][k]
						k=0
边界条件:
	dfs(m−1,j)=grid[m−1][j]
*/

/*
// 会超时的递归代码
func minPathCost(grid [][]int, moveCost [][]int) int {
    m, n := len(grid), len(grid[0])
    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i == m-1 { // 递归边界
            return grid[i][j]
        }
        res := math.MaxInt
        for k, c := range moveCost[grid[i][j]] { // 移动到下一行的第 k 列
            res = min(res, dfs(i+1, k)+c)
        }
        return res + grid[i][j]
    }
    ans := math.MaxInt
    for j := 0; j < n; j++ { // 枚举起点
        ans = min(ans, dfs(0, j))
    }
    return ans
}
*/

/*
记忆化搜索

	func minPathCost(grid [][]int, moveCost [][]int) int {
	    m, n := len(grid), len(grid[0])
	    memo := make([][]int, m)
	    for i := range memo {
	        memo[i] = make([]int, n)
	    }
	    var dfs func(int, int) int
	    dfs = func(i, j int) int {
	        if i == m-1 { // 递归边界
	            return grid[i][j]
	        }
	        p := &memo[i][j]
	        if *p != 0 { // 之前计算过
	            return *p
	        }
	        res := math.MaxInt
	        for k, c := range moveCost[grid[i][j]] { // 移动到下一行的第 k 列
	            res = min(res, dfs(i+1, k)+c)
	        }
	        *p = res + grid[i][j] // 记忆化
	        return *p
	    }
	    ans := math.MaxInt
	    for j := 0; j < n; j++ { // 枚举起点
	        ans = min(ans, dfs(0, j))
	    }
	    return ans
	}
*/

/*
递推

	func minPathCost(grid [][]int, moveCost [][]int) int {
	    m, n := len(grid), len(grid[0])
	    f := make([][]int, m)
	    for i := range f {
	        f[i] = make([]int, n)
	    }
	    f[m-1] = grid[m-1]
	    for i := m - 2; i >= 0; i-- {
	        for j, g := range grid[i] {
	            f[i][j] = math.MaxInt
	            for k, c := range moveCost[g] { // 移动到下一行的第 k 列
	                f[i][j] = min(f[i][j], f[i+1][k]+c)
	            }
	            f[i][j] += g
	        }
	    }
	    return slices.Min(f[0])
	}
*/
func main() {

}
