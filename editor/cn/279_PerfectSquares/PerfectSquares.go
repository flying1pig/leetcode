package main

import "math"

/**
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



 示例 1：


输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

 示例 2：


输入：n = 13
输出：2
解释：13 = 4 + 9



 提示：


 1 <= n <= 10⁴


 Related Topics 广度优先搜索 数学 动态规划 👍 2145 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 写在外面，多个测试数据之间可以共享，减少计算量
var memo [101][10001]int

func init() {
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
}

func dfs(i, j int) int {
	if i == 0 {
		if j == 0 {
			return 0
		}
		return math.MaxInt
	}
	p := &memo[i][j]
	if *p != -1 { // 之前计算过
		return *p
	}
	if j < i*i {
		*p = dfs(i-1, j) // 只能不选
	} else {
		*p = min(dfs(i-1, j), dfs(i, j-i*i)+1) // 不选 vs 选
	}
	return *p
}

func numSquares(n int) int {
	return dfs(int(math.Sqrt(float64(n))), n)
}

//时间复杂度: o(n*根号n)
//空间复杂度: o(n*根号n)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
