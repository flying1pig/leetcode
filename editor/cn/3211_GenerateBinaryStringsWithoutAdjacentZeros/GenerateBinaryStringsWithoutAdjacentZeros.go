package main

/**
给你一个正整数 n。

 如果一个二进制字符串 x 的所有长度为 2 的子字符串中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。

 返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。



 示例 1：


 输入： n = 3


 输出： ["010","011","101","110","111"]

 解释：

 长度为 3 的有效字符串有："010"、"011"、"101"、"110" 和 "111"。

 示例 2：


 输入： n = 1


 输出： ["0","1"]

 解释：

 长度为 1 的有效字符串有："0" 和 "1"。



 提示：


 1 <= n <= 18


 Related Topics 位运算 字符串 回溯 👍 40 👎 0

*/

/*
题型: 回溯
题目: 生成不含相邻零的二进制字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func validStrings(n int) (ans []string) {
	path := make([]byte, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path)) // 注意 string(path) 需要 O(n) 时间
			return
		}

		// 填 1
		path[i] = '1'
		dfs(i + 1)

		// 填 0
		if i == 0 || path[i-1] == '1' {
			path[i] = '0' // 直接覆盖
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
