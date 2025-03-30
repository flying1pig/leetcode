package main

import "strconv"

/**
给你一个二进制字符串 s ，你需要将字符串分割成一个或者多个 子字符串 ，使每个子字符串都是 美丽 的。

 如果一个字符串满足以下条件，我们称它是 美丽 的：


 它不包含前导 0 。
 它是 5 的幂的 二进制 表示。


 请你返回分割后的子字符串的 最少 数目。如果无法将字符串 s 分割成美丽子字符串，请你返回 -1 。

 子字符串是一个字符串中一段连续的字符序列。



 示例 1：

 输入：s = "1011"
输出：2
解释：我们可以将输入字符串分成 ["101", "1"] 。
- 字符串 "101" 不包含前导 0 ，且它是整数 5¹ = 5 的二进制表示。
- 字符串 "1" 不包含前导 0 ，且它是整数 5⁰ = 1 的二进制表示。
最少可以将 s 分成 2 个美丽子字符串。


 示例 2：

 输入：s = "111"
输出：3
解释：我们可以将输入字符串分成 ["1", "1", "1"] 。
- 字符串 "1" 不包含前导 0 ，且它是整数 5⁰ = 1 的二进制表示。
最少可以将 s 分成 3 个美丽子字符串。


 示例 3：

 输入：s = "0"
输出：-1
解释：无法将给定字符串分成任何美丽子字符串。




 提示：


 1 <= s.length <= 15
 s[i] 要么是 '0' 要么是 '1' 。


 Related Topics 哈希表 字符串 动态规划 回溯 👍 19 👎 0

*/

/*
题型: dp
题目: 将字符串分割为最少的美丽子字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
var pow5 []string

func init() {
	// 预处理 2**15 以内的 5 的幂
	for p5 := 1; p5 < 1<<15; p5 *= 5 {
		pow5 = append(pow5, strconv.FormatUint(uint64(p5), 2))
	}
}

func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i] = n + 1
		if s[i] == '0' {
			continue
		}
		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if s[i:i+len(t)] == t {
				f[i] = min(f[i], f[i+len(t)]+1)
			}
		}
	}
	if f[0] > n {
		return -1
	}
	return f[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
