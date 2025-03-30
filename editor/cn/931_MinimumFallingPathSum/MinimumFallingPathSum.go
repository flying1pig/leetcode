package main

import (
	"fmt"
	"math"
)

/**
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个
元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1,
col + 1) 。



 示例 1：




输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
输出：13
解释：如图所示，为和最小的两条下降路径


 示例 2：




输入：matrix = [[-19,57],[-40,-5]]
输出：-59
解释：如图所示，为和最小的下降路径




 提示：


 n == matrix.length == matrix[i].length
 1 <= n <= 100
 -100 <= matrix[i][j] <= 100


 👍 393 👎 0

*/

/*
题型: dp
题目: 下降路径最小和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {
	var f func(int, int) int
	n := len(matrix)
	mem := make([][]int, n)
	mem[0] = matrix[0]
	for i := 1; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = math.MaxInt
		}
	}

	f = func(i int, j int) int {
		if i < 0 || j < 0 || j >= n {
			return math.MaxInt
		}
		if i == 0 {
			return mem[0][j]
		}
		if mem[i][j] == math.MaxInt {
			mem[i][j] = min(f(i-1, j-1), f(i-1, j), f(i-1, j+1)) + matrix[i][j]
		}
		return mem[i][j]
	}

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, f(n-1, i))
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)为下降到(i,j)的路径最小和
状态方程:
	f(i,j) = min(f(i-1,j-1),f(i-1,j),f(i-1,j+1))+matrix[i][j]
边界条件:
	f(-1,j) = f(i,-1) = f(i,n) = math.MaxInt
*/

/*
记忆化搜索:

	func minFallingPathSum(matrix [][]int) int {
	    var f func(int,int) int
		n := len(matrix)
		mem := make([][]int,n)
		mem[0] = matrix[0]
		for i:=1;i<n;i++ {
			mem[i] = make([]int,n)
			for j:=0;j<n;j++ {
				mem[i][j] = math.MaxInt
			}
		}

		f = func(i int, j int) int {
			if i<0 || j<0 || j>=n {
				return math.MaxInt
			}
			if i == 0 {
				return mem[0][j]
			}
			if mem[i][j] == math.MaxInt {
				mem[i][j] = min(f(i-1,j-1),f(i-1,j),f(i-1,j+1)) + matrix[i][j]
			}
			return mem[i][j]
		}

		ans := math.MaxInt
		for i := 0;i<n;i++ {
			ans = min(ans,f(n-1,i))
		}
		return ans
	}

时间复杂度: o(n^2)
空间复杂度: o(n^2)
*/
func main() {
	matrix := [][]int{
		[]int{2, 1, 3},
		[]int{6, 5, 4},
		[]int{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}
