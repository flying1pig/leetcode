package main

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

 问总共有多少条不同的路径？



 示例 1：


输入：m = 3, n = 7
输出：28

 示例 2：


输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下


 示例 3：


输入：m = 7, n = 3
输出：28


 示例 4：


输入：m = 3, n = 3
输出：6



 提示：


 1 <= m, n <= 100
 题目数据保证答案小于等于 2 * 10⁹


 👍 2202 👎 0

*/

/*
题型: dp
题目: 不同路径
*/

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	mem := make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				mem[1] = 1
				continue
			}
			mem[j+1] = mem[j+1] + mem[j]
		}
	}
	return mem[n]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)为走到(i,j)的不同路径数量
	从上面来，f(i,j) = f(i-1,j)
	从左边来，f(i,j) = f(i,j-1)
状态方程: f(i,j) = f(i-1,j) + f(i,j-1)
边界条件:
	f(0,0) = 1
	f(-1,j) = 0
	f(i,-1) = 0
*/

/*
递归:
func uniquePaths(m int, n int) int {
	var f func(i,j int) int
	f = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i<0 || j <0 {
			return 0
		}
		return f(i-1,j) + f(i,j-1)
	}
	return f(m-1,n-1)
}
时间复杂度: o(2^m*n)
空间复杂度: o(m*n)
*/

/*
记忆化搜索:
func uniquePaths(m int, n int) int {
	var f func(i,j int) int
	mem := make([][]int,m)
	for i:=0;i<m;i++ {
		mem[i] = make([]int,n)
	}
	f = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i<0 || j <0 {
			return 0
		}
		if mem[i][j] == 0 {
			mem[i][j] = f(i-1,j) + f(i,j-1)
		}
		return mem[i][j]
	}
	return f(m-1,n-1)
}
时间复杂度：o(m*n)
空间复杂度: o(m*n)
*/

/*
递归:
func uniquePaths(m int, n int) int {
	mem := make([][]int,m+1)
	for i:=0;i<=m;i++ {
		mem[i] = make([]int,n+1)
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if i == 0 && j == 0 {
				mem[1][1] = 1
				continue
			}
			mem[i+1][j+1] = mem[i][j+1] + mem[i+1][j]
		}
	}
	return mem[m][n]
}
时间复杂度: o(m*n)
空间复杂度: o(m*n)
*/

/*
优化:

	func uniquePaths(m int, n int) int {
		mem := make([]int,n+1)
		for i:=0;i<m;i++ {
			for j:=0;j<n;j++ {
				if i == 0 && j == 0 {
					mem[1] = 1
					continue
				}
				mem[j+1] = mem[j+1] + mem[j]
			}
		}
		return mem[n]
	}

时间复杂度: o(m*n)
空间复杂度: o(n)
*/
func main() {}
