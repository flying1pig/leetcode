package main

/**
给你一个字符串 s 和一个整数 k 。请你用 s 字符串中 所有字符 构造 k 个非空 回文串 。

 如果你可以用 s 中所有字符构造 k 个回文字符串，那么请你返回 True ，否则返回 False 。



 示例 1：


输入：s = "annabelle", k = 2
输出：true
解释：可以用 s 中所有字符构造 2 个回文字符串。
一些可行的构造方案包括："anna" + "elble"，"anbna" + "elle"，"anellena" + "b"


 示例 2：


输入：s = "leetcode", k = 3
输出：false
解释：无法用 s 中所有字符构造 3 个回文串。


 示例 3：


输入：s = "true", k = 4
输出：true
解释：唯一可行的方案是让 s 中每个字符单独构成一个字符串。




 提示：


 1 <= s.length <= 10⁵
 s 中所有字符都是小写英文字母。
 1 <= k <= 10⁵


 Related Topics 贪心 哈希表 字符串 计数 👍 57 👎 0

*/

/*
题型: 字符串贪心
题目: 构造 K 个回文字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(s string, k int) bool {
	n := len(s)
	if k > n {
		return false
	}
	dict := map[byte]int{}
	for i := 0; i < n; i++ {
		dict[s[i]]++
	}
	cnt := 0
	for _, v := range dict {
		if v&1 == 1 {
			cnt++
		}
	}
	return cnt <= k
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
