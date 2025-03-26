package main

import "math"

/**
给你两个 正 整数 n 和 x 。

 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。换句话说，你需要返回互不相同整数 [n1, n2, ..., nk] 的集合数目，满足
n = n1ˣ + n2ˣ + ... + nkˣ 。

 由于答案可能非常大，请你将它对 10⁹ + 7 取余后返回。

 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 2³ + 3³ + 5³ 。



 示例 1：

 输入：n = 10, x = 2
输出：1
解释：我们可以将 n 表示为：n = 3² + 1² = 10 。
这是唯一将 10 表达成不同整数 2 次方之和的方案。


 示例 2：

 输入：n = 4, x = 1
输出：2
解释：我们可以将 n 按以下方案表示：
- n = 4¹ = 4 。
- n = 3¹ + 1¹ = 4 。




 提示：


 1 <= n <= 300
 1 <= x <= 5


 Related Topics 动态规划 👍 31 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfWays(n int, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for c := n; c >= v; c-- {
			f[c] += f[c-v]
		}
	}
	return f[n] % 1_000_000_007
}

func pow(i, x int) int {
	return int(math.Pow(float64(i), float64(x)))
}

//leetcode submit region end(Prohibit modification and deletion)

/*
把n看成背包容量，1^x,2^x,3^x...看成物品，本题就是0-1背包模板题
*/
func main() {

}
