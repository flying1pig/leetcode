package main

import (
	"math/bits"
	"slices"
)

/**
给你一个大小为 m x n 的二维整数数组 grid 和一个整数 k 。

 你的任务是统计满足以下 条件 且从左上格子 (0, 0) 出发到达右下格子 (m - 1, n - 1) 的路径数目：


 每一步你可以向右或者向下走，也就是如果格子存在的话，可以从格子 (i, j) 走到格子 (i, j + 1) 或者格子 (i + 1, j) 。
 路径上经过的所有数字 XOR 异或值必须 等于 k 。


 请你返回满足上述条件的路径总数。

 由于答案可能很大，请你将答案对 10⁹ + 7 取余 后返回。



 示例 1：


 输入：grid = [[2, 1, 5], [7, 10, 0], [12, 6, 4]], k = 11


 输出：3

 解释：

 3 条路径分别为：


 (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)
 (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
 (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)


 示例 2：


 输入：grid = [[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]], k = 2


 输出：5

 解释：

 5 条路径分别为：


 (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)
 (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)
 (0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)
 (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)
 (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)


 示例 3：


 输入：grid = [[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]], k = 10


 输出：0



 提示：


 1 <= m == grid.length <= 300
 1 <= n == grid[r].length <= 300
 0 <= grid[r][c] < 16
 0 <= k < 16


 👍 4 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func countPathsWithXorValue(grid [][]int, k int) int {
	mod := 1_000_000_007
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	u := 1 << bits.Len(uint(mx))
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, u)
			for x := range memo[i][j] {
				memo[i][j][x] = -1
			}
		}
	}

	var f func(int, int, int) int
	f = func(i int, j int, x int) int {
		if i < 0 || j < 0 {
			return 0
		}
		val := grid[i][j]
		if i == 0 && j == 0 {
			if val == x {
				return 1
			}
			return 0
		}
		if memo[i][j][x] == -1 {
			memo[i][j][x] = (f(i, j-1, x^val) + f(i-1, j, x^val)) % mod
		}
		return memo[i][j][x]
	}
	return f(m-1, n-1, k)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j,x)表示从(0,0)到当前位置(i,j)且xor值为x的方案数。
状态方程:
	f(i,j,x) = f(i-1,j,x^grid[i][j])+f(i,j-1,x^grid[i][j])
	^表示异或运算
边界条件:
	f(-1,j,x) = f(i,-1,x) = 0
	f(0,0,grid[0][0]) = 1, 其余f(0,0,x) = 0
细节:
	设mx为grid[i][j]的最大值，其二进制长度为L，那么(2的L次方-1)就是XOR能取到的最大值。
	如果k> (2的L次方-1)，那么直接返回0
*/

/*
记忆化搜索:

	func countPathsWithXorValue(grid [][]int, k int) int {
	    mod := 1_000_000_007
		mx := 0
		for _,row := range grid {
			mx = max(mx,slices.Max(row))
		}
		u := 1<<bits.Len(uint(mx))
		if k>=u {
			return 0
		}

		m, n := len(grid), len(grid[0])
		memo := make([][][]int, m)
		for i := range memo {
			memo[i] = make([][]int, n)
			for j := range memo[i] {
				memo[i][j] = make([]int, u)
				for x := range memo[i][j] {
					memo[i][j][x] = -1
				}
			}
		}

		var f func(int,int,int) int
		f = func(i int, j int, x int) int {
			if i<0|| j<0 {
				return 0
			}
			val := grid[i][j]
			if i == 0 && j ==0 {
				if val == x {
					return 1
				}
				return 0
			}
			if memo[i][j][x] == -1 {
				memo[i][j][x] = (f(i,j-1,x^val)+f(i-1,j,x^val))%mod
			}
			return memo[i][j][x]
		}
		return f(m-1,n-1,k)
	}

时间复杂度: o(mnU),U是grid最大值
空间复杂度: o(mnU)
*/
func main() {

}
