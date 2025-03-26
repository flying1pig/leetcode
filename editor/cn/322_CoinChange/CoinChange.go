package main

import "math"

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

 你可以认为每种硬币的数量是无限的。



 示例 1：


输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

 示例 2：


输入：coins = [2], amount = 3
输出：-1

 示例 3：


输入：coins = [1], amount = 0
输出：0




 提示：


 1 <= coins.length <= 12
 1 <= coins[i] <= 2³¹ - 1
 0 <= amount <= 10⁴


 Related Topics 广度优先搜索 数组 动态规划 👍 3017 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有访问过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2 // 除 2 防止下面 + 1 溢出
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
