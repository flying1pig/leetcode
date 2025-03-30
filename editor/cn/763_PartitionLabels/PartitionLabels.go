package main

/**
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab",
"cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。

 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

 返回一个表示每个字符串片段的长度的列表。


示例 1：


输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。

 示例 2：


输入：s = "eccbbbbdec"
输出：[10]




 提示：


 1 <= s.length <= 500
 s 仅由小写英文字母组成


 Related Topics 贪心 哈希表 双指针 字符串 👍 1267 👎 0

*/

/*
题型: 贪心
题目: 划分字母区间
*/

// leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) (ans []int) {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i // 每个字母最后出现的下标
	}

	start, end := 0, 0
	for i, c := range s {
		end = max(end, last[c-'a']) // 更新当前区间右端点的最大值
		if end == i {               // 当前区间合并完毕
			ans = append(ans, end-start+1) // 区间长度加入答案
			start = i + 1                  // 下一个区间的左端点
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
