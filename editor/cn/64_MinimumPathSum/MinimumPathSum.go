package main

import "math"

/**
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

 说明：每次只能向下或者向右移动一步。



 示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。


 示例 2：


输入：grid = [[1,2,3],[4,5,6]]
输出：12




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 200
 0 <= grid[i][j] <= 200


 👍 1794 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	n := len(grid[0])
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = math.MaxInt
	}

	mem[1] = 0
	for _, row := range grid {
		for i, x := range row {
			mem[i+1] = min(mem[i], mem[i+1]) + x
		}
	}
	return mem[n]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)为从(0,0)走到(i,j)的最小路径和
如果从左边过来，f(i,j) = f(i,j-1) + grid[i][j]
如果从上边过来，f(i,j) = f(i-1,j) + grid[i][j]
状态方程:
	f(i,j) = min(f(i,j-1),f(i-1,j)) + grid[i][j]
边界条件:
	f(0,0) = grid[0][0]
	f(-1,j) = f (i,-1) = math.MaxInt
*/

/*
递归:
func minPathSum(grid [][]int) int {
    var f func(i,j int) int
	f = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		return min(f(i,j-1),f(i-1,j)) + grid[i][j]
	}
	return f(len(grid)-1,len(grid[0])-1)
}
时间复杂度: o(2^m*n)
空间复杂度: o(m+n)
*/

/*
记忆优化搜索:
func minPathSum(grid [][]int) int {
    var f func(i,j int) int
	m,n := len(grid),len(grid[0])
	mem := make([][]int,m)
	for i := range mem {
		mem[i] = make([]int,n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	f = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		mem[i][j] = min(f(i,j-1),f(i-1,j)) + grid[i][j]
		return mem[i][j]
	}
	return f(len(grid)-1,len(grid[0])-1)
}
时间复杂度: o(m*n)
空间复杂度: o(m*n)
*/

/*
递推:
递推过程中为了不使数组下标越界，即把f(i,-1)和f(-1,j)这样的状态也翻译过来，状态方程可以这样写:
f(i+1,j+1) = min(f(i+1,j),f(i,j+1))+grid[i][j]

func minPathSum(grid [][]int) int {
	m,n := len(grid),len(grid[0])
	mem := make([][]int,m+1)
	for i := range mem {
		mem[i] = make([]int,n+1)
		for j := range mem[i] {
			mem[i][j] = math.MaxInt
		}
	}

	for i,row := range grid {
		for j,x := range row {
			if i ==0 && j == 0 {
				mem[1][1] = x
			} else {
				mem[i+1][j+1] = min(mem[i+1][j],mem[i][j+1]) + x
			}
		}
	}
	return mem[m][n]
}
时间复杂度: o(m*n)
空间复杂度: o(m*n)
*/

/*
空间优化:
举个例子, 在计算f(1,1)时，会用到f(0,1)，但是之后就不再用到。那么干脆把f(1,1)记到f(0,1)中，这样对于f(1,2)来说，
它需要的数据就在f(0,1)和f(0,2)中。f(1,2)算完后也可以记录到f(0,2)中。所以只需要维护一个长度为n+1的数据就够了。

func minPathSum(grid [][]int) int {
	n := len(grid[0])
	mem := make([]int,n+1)
	for i := range mem {
		mem[i] = math.MaxInt
	}

	mem[1] = 0
	for _,row := range grid {
		for i,x := range row {
			mem[i+1] = min(mem[i],mem[i+1])+x
		}
	}
	return mem[n]
}
时间复杂度: o(mn)
空间复杂度: o(n)
*/
func main() {

}
