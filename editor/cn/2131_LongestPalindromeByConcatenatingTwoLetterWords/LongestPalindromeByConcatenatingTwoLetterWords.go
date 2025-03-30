package main

/**
给你一个字符串数组 words 。words 中每个元素都是一个包含 两个 小写英文字母的单词。

 请你从 words 中选择一些元素并按 任意顺序 连接它们，并得到一个 尽可能长的回文串 。每个元素 至多 只能使用一次。

 请你返回你能得到的最长回文串的 长度 。如果没办法得到任何一个回文串，请你返回 0 。

 回文串 指的是从前往后和从后往前读一样的字符串。



 示例 1：

 输入：words = ["lc","cl","gg"]
输出：6
解释：一个最长的回文串为 "lc" + "gg" + "cl" = "lcggcl" ，长度为 6 。
"clgglc" 是另一个可以得到的最长回文串。


 示例 2：

 输入：words = ["ab","ty","yt","lc","cl","ab"]
输出：8
解释：最长回文串是 "ty" + "lc" + "cl" + "yt" = "tylcclyt" ，长度为 8 。
"lcyttycl" 是另一个可以得到的最长回文串。


 示例 3：

 输入：words = ["cc","ll","xx"]
输出：2
解释：最长回文串是 "cc" ，长度为 2 。
"ll" 是另一个可以得到的最长回文串。"xx" 也是。



 提示：


 1 <= words.length <= 10⁵
 words[i].length == 2
 words[i] 仅包含小写英文字母。


 Related Topics 贪心 数组 哈希表 字符串 计数 👍 58 👎 0

*/

/*
题型: 字符串贪心
题目: 连接两字母单词得到的最长回文串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(words []string) (ans int) {
	cnt := [26][26]int{}
	for _, s := range words {
		cnt[s[0]-'a'][s[1]-'a']++
	}
	odd := 0 // 是否有一个出现奇数次的 AA
	for i := 0; i < 26; i++ {
		c := cnt[i][i] // 相同字符
		ans += c &^ 1  // c &^ 1 等价于 c - c % 2
		odd |= c & 1
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2 // AB BA 取出现次数最小值
		}
	}
	return (ans + odd) * 2 // 上面的循环统计的是字符串个数，最后再乘 2
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
