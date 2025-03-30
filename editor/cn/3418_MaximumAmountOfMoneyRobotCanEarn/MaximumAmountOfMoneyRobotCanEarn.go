package main

import "math"

/**
给你一个 m x n 的网格。一个机器人从网格的左上角 (0, 0) 出发，目标是到达网格的右下角 (m - 1, n - 1)。在任意时刻，机器人只能向右或向
下移动。

 网格中的每个单元格包含一个值 coins[i][j]：


 如果 coins[i][j] >= 0，机器人可以获得该单元格的金币。
 如果 coins[i][j] < 0，机器人会遇到一个强盗，强盗会抢走该单元格数值的 绝对值 的金币。


 机器人有一项特殊能力，可以在行程中 最多感化 2个单元格的强盗，从而防止这些单元格的金币被抢走。

 注意：机器人的总金币数可以是负数。

 返回机器人在路径上可以获得的 最大金币数 。



 示例 1：


 输入： coins = [[0,1,-1],[1,-2,3],[2,-3,4]]


 输出： 8

 解释：

 一个获得最多金币的最优路径如下：


 从 (0, 0) 出发，初始金币为 0（总金币 = 0）。
 移动到 (0, 1)，获得 1 枚金币（总金币 = 0 + 1 = 1）。
 移动到 (1, 1)，遇到强盗抢走 2 枚金币。机器人在此处使用一次感化能力，避免被抢（总金币 = 1）。
 移动到 (1, 2)，获得 3 枚金币（总金币 = 1 + 3 = 4）。
 移动到 (2, 2)，获得 4 枚金币（总金币 = 4 + 4 = 8）。


 示例 2：


 输入： coins = [[10,10,10],[10,10,10]]


 输出： 40

 解释：

 一个获得最多金币的最优路径如下：


 从 (0, 0) 出发，初始金币为 10（总金币 = 10）。
 移动到 (0, 1)，获得 10 枚金币（总金币 = 10 + 10 = 20）。
 移动到 (0, 2)，再获得 10 枚金币（总金币 = 20 + 10 = 30）。
 移动到 (1, 2)，获得 10 枚金币（总金币 = 30 + 10 = 40）。




 提示：


 m == coins.length
 n == coins[i].length
 1 <= m, n <= 500
 -1000 <= coins[i][j] <= 1000


 Related Topics 数组 动态规划 矩阵 👍 9 👎 0

*/

/*
题型: dp
题目: 机器人可以获得的最大金币数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][3]int, m)
	for i := range memo {
		memo[i] = make([][3]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		x := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return x
			}
			return max(x, 0)
		}
		p := &memo[i][j][k]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x // 选
		if x < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1)) // 不选
		}
		*p = res // 记忆化
		return res
	}
	return dfs(m-1, n-1, 2)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
	选: dfs(i,j,k)=max(dfs(i−1,j,k),dfs(i,j−1,k))+coins[i][j]
	不选: dfs(i,j,k)=max(dfs(i−1,j,k−1),dfs(i,j−1,k−1))   k>0 且 coins[i][j]<0
边界条件:
	dfs(−1,j,k)=dfs(i,−1,k)= -inf
	dfs(0,0,0)=coins[0][0]
	dfs(0,0,k>0)=max(coins[0][0],0)
*/

/*
记忆化搜索:
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][3]int, m)
	for i := range memo {
		memo[i] = make([][3]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		x := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return x
			}
			return max(x, 0)
		}
		p := &memo[i][j][k]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x // 选
		if x < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1)) // 不选
		}
		*p = res // 记忆化
		return res
	}
	return dfs(m-1, n-1, 2)
}
*/

/*
递推:
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	f := make([][][3]int, m+1)
	for i := range f {
		f[i] = make([][3]int, n+1)
	}
	for j := range f[0] {
		f[0][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[0][1] = [3]int{}
	for i, row := range coins {
		f[i+1][0] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		for j, x := range row {
			f[i+1][j+1][0] = max(f[i+1][j][0], f[i][j+1][0]) + x
			f[i+1][j+1][1] = max(f[i+1][j][1]+x, f[i][j+1][1]+x, f[i+1][j][0], f[i][j+1][0])
			f[i+1][j+1][2] = max(f[i+1][j][2]+x, f[i][j+1][2]+x, f[i+1][j][1], f[i][j+1][1])
		}
	}
	return f[m][n][2]
}
*/

/*
空间优化

	func maximumAmount(coins [][]int) int {
		n := len(coins[0])
		f := make([][3]int, n+1)
		for j := range f {
			f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		}
		f[1] = [3]int{}
		for _, row := range coins {
			for j, x := range row {
				f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x, f[j][1], f[j+1][1])
				f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x, f[j][0], f[j+1][0])
				f[j+1][0] = max(f[j][0], f[j+1][0]) + x
			}
		}
		return f[n][2]
	}
*/
func main() {

}
