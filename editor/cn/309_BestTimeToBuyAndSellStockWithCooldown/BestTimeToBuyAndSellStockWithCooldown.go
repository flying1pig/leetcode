package main

import "math"

/**
给定一个整数数组
 prices，其中第 prices[i] 表示第 i 天的股票价格 。

 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:


 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。


 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



 示例 1:


输入: prices = [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

 示例 2:


输入: prices = [1]
输出: 0




 提示：


 1 <= prices.length <= 5000
 0 <= prices[i] <= 1000


 Related Topics 数组 动态规划 👍 1836 👎 0

*/

/*
题型: dp
题目: 买卖股票的最佳时机含冷冻期
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	pre0, f0, f1 := 0, 0, math.MinInt
	for _, p := range prices {
		pre0, f0, f1 = f0, max(f0, f1+p), max(f1, pre0-p)
	}
	return f0
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义dfs(i,0)表示第i天结束时，未持有股票的最大利润
定义dfs(i,1)表示第i天结束时，持有股票的最大利润
状态方程:
	dfs(i,0) = max(dfs(i-1,0),dfs(i-1,1)+prices[i])
	dfs(i,1) = max(dfs(i-1,1),dfs(i-2,0)-prices[i])
边界条件:
	dfs(-1,0) = 0
	dfs(-2,0) = 0
	dfs(-1,1) = math.MinInt
*/

/*
递推:
func maxProfit(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return
		}
		p := &memo[i][hold]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0)
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func maxProfit(prices []int) int {
    n := len(prices)
    f := make([][2]int, n+2)
    f[1][1] = math.MinInt
    for i, p := range prices {
        f[i+2][0] = max(f[i+1][0], f[i+1][1]+p)
        f[i+2][1] = max(f[i+1][1], f[i][0]-p)
    }
    return f[n+1][0]
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
空间优化:

	func maxProfit(prices []int) int {
	    pre0, f0, f1 := 0, 0, math.MinInt
	    for _, p := range prices {
	        pre0, f0, f1 = f0, max(f0, f1+p), max(f1, pre0-p)
	    }
	    return f0
	}

时间复杂度: o(n)
空间复杂度: o(1)
*/
func main() {

}
