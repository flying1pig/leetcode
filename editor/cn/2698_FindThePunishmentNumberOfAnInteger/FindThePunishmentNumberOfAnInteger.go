package main

import "strconv"

/**
给你一个正整数 n ，请你返回 n 的 惩罚数 。

 n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：


 1 <= i <= n
 i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。




 示例 1：


输入：n = 10
输出：182
解释：总共有 3 个范围在 [1, 10] 的整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
因此，10 的惩罚数为 1 + 81 + 100 = 182


 示例 2：


输入：n = 37
输出：1478
解释：总共有 4 个范围在 [1, 37] 的整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
- 36 ，因为 36 * 36 = 1296 ，且 1296 可以分割成 1 + 29 + 6 。
因此，37 的惩罚数为 1 + 81 + 100 + 1296 = 1478




 提示：


 1 <= n <= 1000


 Related Topics 数学 回溯 👍 141 👎 0

*/

/*
题型: 回溯
题目: 求一个整数的惩罚数
*/

// leetcode submit region begin(Prohibit modification and deletion)
var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 符合要求
			}
			x := 0
			for j := p; j < n; j++ { // 枚举分割出从 s[p] 到 s[j] 的子串
				x = x*10 + int(s[j]-'0') // 子串对应的整数值
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
