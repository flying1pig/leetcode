package main

import "math"

/**
给你一个 n x n 的二进制网格 grid，每一次操作中，你可以选择网格的 相邻两行 进行交换。

 一个符合要求的网格需要满足主对角线以上的格子全部都是 0 。

 请你返回使网格满足要求的最少操作次数，如果无法使网格符合要求，请你返回 -1 。

 主对角线指的是从 (1, 1) 到 (n, n) 的这些格子。



 示例 1：



 输入：grid = [[0,0,1],[1,1,0],[1,0,0]]
输出：3


 示例 2：



 输入：grid = [[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]
输出：-1
解释：所有行都是一样的，交换相邻行无法使网格符合要求。


 示例 3：



 输入：grid = [[1,0,0],[1,1,0],[1,1,1]]
输出：0




 提示：


 n == grid.length
 n == grid[i].length
 1 <= n <= 200
 grid[i][j] 要么是 0 要么是 1 。


 Related Topics 贪心 数组 矩阵 👍 57 👎 0

*/

/*
题型: 贪心
题目: 排布二进制网格的最少交换次数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minSwaps(grid [][]int) int {
	n := len(grid)
	//前缀和,用于后继快速计算是否符合本行要求
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += grid[i][j-1]
		}
	}

	arr := make([][2]int, n)
	var dfs func(i int, arr [][2]int) (ans int)
	dfs = func(i int, arr [][2]int) (ans int) {
		if i == n-1 {
			return 0
		}

		ans = math.MaxInt32 / 2
		for j := 0; j < n; j++ {
			if arr[j][1] == 1 {
				continue
			}
			//贪心 优先使用最早能满足主对角线的行
			if grid[j][n-1]-grid[j][i] == 0 { //i行主对角线右侧需要全部是0
				arr[j][1] = 1                            //已经使用
				return dfs(i+1, arr) + j - i + arr[j][0] //j - i上浮的行数
			}
			arr[j][0]++ //当轮没有被选中的行,必然要下沉一行
		}

		return
	}
	ans := dfs(0, arr)
	if ans >= math.MaxInt32/2 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
