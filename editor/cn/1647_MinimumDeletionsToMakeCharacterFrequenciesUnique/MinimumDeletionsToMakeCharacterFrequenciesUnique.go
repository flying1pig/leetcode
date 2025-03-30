package main

/**
如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。

 给你一个字符串 s，返回使 s 成为 优质字符串 需要删除的 最小 字符数。

 字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，'a' 的频次是 2，而 'b' 的频次是 1 。



 示例 1：


输入：s = "aab"
输出：0
解释：s 已经是优质字符串。


 示例 2：


输入：s = "aaabbbcc"
输出：2
解释：可以删除两个 'b' , 得到优质字符串 "aaabcc" 。
另一种方式是删除一个 'b' 和一个 'c' ，得到优质字符串 "aaabbc" 。

 示例 3：


输入：s = "ceabaacb"
输出：2
解释：可以删除两个 'c' 得到优质字符串 "eabaab" 。
注意，只需要关注结果字符串中仍然存在的字符。（即，频次为 0 的字符会忽略不计。）




 提示：


 1 <= s.length <= 10⁵
 s 仅含小写英文字母


 Related Topics 贪心 哈希表 字符串 排序 👍 79 👎 0

*/

/*
题型: 贪心
题目: 字符频次唯一的最小删除次数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minDeletions(s string) (ans int) {
	// 初始化 freq 数组
	freq := make([]int, 26)

	// 统计每个字符出现的频次
	for i := range s {
		freq[s[i]-'a']++
	}

	// 使用哈希表记录出现过的频次
	seen := make(map[int]bool)

	for _, f := range freq {
		if f == 0 {
			continue
		}
		for f > 0 && seen[f] {
			f--
			ans++
		}
		seen[f] = true
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
