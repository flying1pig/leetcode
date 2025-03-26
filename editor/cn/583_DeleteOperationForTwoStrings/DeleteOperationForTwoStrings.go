package main

/**
给定两个单词 word1 和
 word2 ，返回使得
 word1 和
 word2 相同所需的最小步数。

 每步 可以删除任意一个字符串中的一个字符。



 示例 1：


输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"


 示例 2:


输入：word1 = "leetcode", word2 = "etco"
输出：4




 提示：



 1 <= word1.length, word2.length <= 500
 word1 和 word2 只包含小写英文字母


 Related Topics 字符串 动态规划 👍 725 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	lcs := dp[m][n]
	return m + n - lcs*2
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路: 先算最长公共子序列长度k，答案为len(word1)+len(word2) - 2*k
func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	lcs := dp[m][n]
	return m + n - lcs*2
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

func main() {

}
