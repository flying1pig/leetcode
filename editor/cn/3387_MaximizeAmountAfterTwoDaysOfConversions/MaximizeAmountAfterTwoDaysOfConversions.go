package main

/**
给你一个字符串 initialCurrency，表示初始货币类型，并且你一开始拥有 1.0 单位的 initialCurrency。

 另给你四个数组，分别表示货币对（字符串）和汇率（实数）：


 pairs1[i] = [startCurrencyi, targetCurrencyi] 表示在 第 1 天，可以按照汇率 rates1[i] 将
startCurrencyi 转换为 targetCurrencyi。
 pairs2[i] = [startCurrencyi, targetCurrencyi] 表示在 第 2 天，可以按照汇率 rates2[i] 将
startCurrencyi 转换为 targetCurrencyi。
 此外，每种 targetCurrency 都可以以汇率 1 / rate 转换回对应的 startCurrency。


 你可以在 第 1 天 使用 rates1 进行任意次数的兑换（包括 0 次），然后在 第 2 天 使用 rates2 再进行任意次数的兑换（包括 0 次）。


 返回在两天兑换后，最大可能拥有的 initialCurrency 的数量。

 注意：汇率是有效的，并且第 1 天和第 2 天的汇率之间相互独立，不会产生矛盾。



 示例 1：


 输入： initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [
2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.
0,6.0]


 输出： 720.00000

 解释：

 根据题目要求，需要最大化最终的 EUR 数量，从 1.0 EUR 开始：


 第 1 天：



 将 EUR 换成 USD，得到 2.0 USD。
 将 USD 换成 JPY，得到 6.0 JPY。


 第 2 天：

 将 JPY 换成 USD，得到 24.0 USD。
 将 USD 换成 CHF，得到 120.0 CHF。
 最后将 CHF 换回 EUR，得到 720.0 EUR。




 示例 2：


 输入： initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 =
 [["NGN","EUR"]], rates2 = [6.0]


 输出： 1.50000

 解释：

 在第 1 天将 NGN 换成 EUR，并在第 2 天用反向汇率将 EUR 换回 NGN，可以最大化最终的 NGN 数量。

 示例 3：


 输入： initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 =
 [["EUR","JPY"]], rates2 = [10.0]


 输出： 1.00000

 解释：

 在这个例子中，不需要在任何一天进行任何兑换。



 提示：


 1 <= initialCurrency.length <= 3
 initialCurrency 仅由大写英文字母组成。
 1 <= n == pairs1.length <= 10
 1 <= m == pairs2.length <= 10
 pairs1[i] == [startCurrencyi, targetCurrencyi]
 pairs2[i] == [startCurrencyi, targetCurrencyi]
 1 <= startCurrencyi.length, targetCurrencyi.length <= 3
 startCurrencyi 和 targetCurrencyi 仅由大写英文字母组成。
 rates1.length == n
 rates2.length == m
 1.0 <= rates1[i], rates2[i] <= 10.0
 输入保证两个转换图在各自的天数中没有矛盾或循环。
 输入保证输出 最大 为 5 * 10¹⁰。


 Related Topics 深度优先搜索 广度优先搜索 图 数组 字符串 👍 7 👎 0

*/

/*
题型: 图论dfs
题目: 两天自由外汇交易后的最大货币数
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	to   string
	rate float64
}

func calcAmount(pairs [][]string, rates []float64, initialCurrency string) map[string]float64 {
	g := map[string][]pair{}
	for i, p := range pairs {
		x, y, r := p[0], p[1], rates[i]
		g[x] = append(g[x], pair{y, r})
		g[y] = append(g[y], pair{x, 1 / r})
	}

	amount := map[string]float64{}
	var dfs func(string, float64)
	dfs = func(x string, curAmount float64) {
		amount[x] = curAmount
		for _, e := range g[x] {
			// 每个节点只需递归一次（重复递归算出来的结果是一样的，因为题目保证汇率没有矛盾）
			if amount[e.to] == 0 {
				dfs(e.to, curAmount*e.rate)
			}
		}
	}
	dfs(initialCurrency, 1)
	return amount
}

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) (ans float64) {
	day1Amount := calcAmount(pairs1, rates1, initialCurrency)
	day2Amount := calcAmount(pairs2, rates2, initialCurrency)
	for x, a2 := range day2Amount {
		ans = max(ans, day1Amount[x]/a2)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
