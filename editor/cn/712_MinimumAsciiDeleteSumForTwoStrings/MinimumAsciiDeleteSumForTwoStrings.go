package main

/**
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。



 示例 1:


输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。


 示例 2:


输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。




 提示:


 0 <= s1.length, s2.length <= 1000
 s1 和 s2 由小写英文字母组成


 Related Topics 字符串 动态规划 👍 423 👎 0

*/

/*
题型: dp
题目: 两个字符串的最小 ASCII 删除和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := range dp[0] {
		if j > 0 {
			dp[0][j] = dp[0][j-1] + int(s2[j-1])
		}
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(c1), dp[i+1][j]+int(c2))
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
创建len(s1)+1行，len(s2)+1列的二维数组, 定义dp[i][j]表示使s1[0:i]和s2[0:j]想通的最小ASCII删除和。
状态方程:
	dp[i][j] = dp[i-1][j-1], s1[i-1] = s2[j-1]
	dp[i][j] = min(dp[i-1][j]+s1[i-1],dp[i][j-1]+s2[j-1]) s1[i-1] != s2[j-1]
边界条件:
	当i=j=0, dp[0][0] = 0
	当i=0且j>0, dp[0][j] = dp[0][j-1] + s2[j-1]
	当j=0且i>0，dp[i][0] = dp[i-1][0]+s1[i-1]
*/

/*
	func minimumDeleteSum(s1 string, s2 string) int {
	    m, n := len(s1), len(s2)
	    dp := make([][]int, m+1)
	    for i := range dp {
	        dp[i] = make([]int, n+1)
	        if i > 0 {
	            dp[i][0] = dp[i-1][0] + int(s1[i-1])
	        }
	    }
	    for j := range dp[0] {
	        if j > 0 {
	            dp[0][j] = dp[0][j-1] + int(s2[j-1])
	        }
	    }
	    for i, c1 := range s1 {
	        for j, c2 := range s2 {
	            if c1 == c2 {
	                dp[i+1][j+1] = dp[i][j]
	            } else {
	                dp[i+1][j+1] = min(dp[i][j+1] + int(c1), dp[i+1][j] + int(c2))
	            }
	        }
	    }
	    return dp[m][n]
	}

	func min(a, b int) int {
	    if a > b {
	        return b
	    }
	    return a
	}

时间复杂度: o(mn)
空间复杂度: o(mn)
*/
func main() {

}
