package main

import "strings"

/**
给你一个仅由数字（0 - 9）组成的字符串 num 。

 请你找出能够使用 num 中数字形成的 最大回文 整数，并以字符串形式返回。该整数不含 前导零 。

 注意：


 你 无需 使用 num 中的所有数字，但你必须使用 至少 一个数字。
 数字可以重新排序。




 示例 1：


输入：num = "444947137"
输出："7449447"
解释：
从 "444947137" 中选用数字 "4449477"，可以形成回文整数 "7449447" 。
可以证明 "7449447" 是能够形成的最大回文整数。


 示例 2：


输入：num = "00009"
输出："9"
解释：
可以证明 "9" 能够形成的最大回文整数。
注意返回的整数不应含前导零。




 提示：


 1 <= num.length <= 10⁵
 num 由数字（0 - 9）组成


 Related Topics 贪心 哈希表 字符串 计数 👍 38 👎 0

*/

/*
题型: 字符串贪心
题目: 最大回文数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestPalindromic(num string) string {
	cnt := ['9' + 1]int{}
	for _, d := range num {
		cnt[d]++
	}
	if cnt['0'] == len(num) { // 特判最特殊的情况：num 全是 0
		return "0"
	}

	s := []byte{}
	for i := byte('9'); i > '0' || i == '0' && len(s) > 0; i-- { // 如果填了数字，则可以填 0
		s = append(s, strings.Repeat(string(i), cnt[i]/2)...)
	}

	j := len(s) - 1
	for i := byte('9'); i >= '0'; i-- {
		if cnt[i]&1 > 0 { // 还可以填一个变成奇回文串
			s = append(s, i)
			break
		}
	}
	for ; j >= 0; j-- { // 添加镜像部分
		s = append(s, s[j])
	}
	return string(s)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
