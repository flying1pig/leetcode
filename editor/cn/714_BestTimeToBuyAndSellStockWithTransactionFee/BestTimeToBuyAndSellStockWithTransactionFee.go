package main

import "math"

/**
给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。

 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

 返回获得利润的最大值。

 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。



 示例 1：


输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
输出：8
解释：能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8

 示例 2：


输入：prices = [1,3,7,5,10,3], fee = 3
输出：6




 提示：


 1 <= prices.length <= 5 * 10⁴
 1 <= prices[i] < 5 * 10⁴
 0 <= fee < 5 * 10⁴


 Related Topics 贪心 数组 动态规划 👍 1128 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2 // 防止溢出
			}
			return
		}
		p := &memo[i][hold]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i]-fee)
	}
	return dfs(n-1, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义dfs(i,0)表示第i天结束时，未持有股票的最大利润
定义dfs(i,1)表示第i天结束时，持有股票的最大利润
状态方程:
	dfs(i,0) = max(dfs(i-1,0),dfs(i-1,1)+prices[i]-fee)
	dfs(i,1) = max(dfs(i-1,1),dfs(i-1,0)-prices[i])
边界条件:
	dfs(-1,0) = 0
	dfs(-1,1) = math.MinInt
*/

/*
递归:

	func maxProfit(prices []int, fee int) int {
		n := len(prices)
		memo := make([][2]int, n)
		for i := range memo {
			memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
		}
		var dfs func(int, int) int
		dfs = func(i, hold int) (res int) {
			if i < 0 {
				if hold == 1 {
					return math.MinInt / 2 // 防止溢出
				}
				return
			}
			p := &memo[i][hold]
			if *p != -1 { // 之前计算过
				return *p
			}
			defer func() { *p = res }() // 记忆化
			if hold == 1 {
				return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
			}
			return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i]-fee)
		}
		return dfs(n-1, 0)
	}

时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:

	func maxProfit(prices []int, fee int) int {
	    n := len(prices)
	    f := make([][2]int, n+1)
	    f[0][1] = math.MinInt / 2
	    for i, p := range prices {
	        f[i+1][0] = max(f[i][0], f[i][1]+p-fee)
	        f[i+1][1] = max(f[i][1], f[i][0]-p)
	    }
	    return f[n][0]
	}

时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
空间优化:

	func maxProfit(prices []int, fee int) int {
	    f0, f1 := 0, math.MinInt/2
	    for _, p := range prices {
	        f0, f1 = max(f0, f1+p-fee), max(f1, f0-p)
	    }
	    return f0
	}

时间复杂度: o(n)
空间复杂度: o(1)
*/
func main() {

}
