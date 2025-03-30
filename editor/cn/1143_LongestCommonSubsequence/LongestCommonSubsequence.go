package main

/**
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。


 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。


 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。



 示例 1：


输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。


 示例 2：


输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。


 示例 3：


输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。




 提示：


 1 <= text1.length, text2.length <= 1000
 text1 和 text2 仅由小写英文字符组成。


 Related Topics 字符串 动态规划 👍 1719 👎 0

*/

/*
题型: dp
题目: 最长公共子序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(s, t string) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if s[i] == t[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(n-1, m-1)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
记忆化搜索:
func longestCommonSubsequence(s, t string) int {
    n, m := len(s), len(t)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if i < 0 || j < 0 {
            return
        }
        p := &memo[i][j]
        if *p != -1 { // 之前计算过
            return *p
        }
        defer func() { *p = res }() // 记忆化
        if s[i] == t[j] {
            return dfs(i-1, j-1) + 1
        }
        return max(dfs(i-1, j), dfs(i, j-1))
    }
    return dfs(n-1, m-1)
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
翻译成递推:
func longestCommonSubsequence(s, t string) int {
    n, m := len(s), len(t)
    f := make([][]int, n+1)
    for i := range f {
        f[i] = make([]int, m+1)
    }
    for i, x := range s {
        for j, y := range t {
            if x == y {
                f[i+1][j+1] = f[i][j] + 1
            } else {
                f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
            }
        }
    }
    return f[n][m]
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
空间优化,两个滚动数组:
func longestCommonSubsequence(s, t string) int {
    n, m := len(s), len(t)
    f := [2][]int{make([]int, m+1), make([]int, m+1)}
    for i, x := range s {
        for j, y := range t {
            if x == y {
                f[(i+1)%2][j+1] = f[i%2][j] + 1
            } else {
                f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j])
            }
        }
    }
    return f[n%2][m]
}
*/

/*
空间优化，一个数组:

	func longestCommonSubsequence(s, t string) int {
	    m := len(t)
	    f := make([]int, m+1)
	    for _, x := range s {
	        pre := 0 // f[0]
	        for j, y := range t {
	            if x == y {
	                f[j+1], pre = pre+1, f[j+1]
	            } else {
	                pre = f[j+1]
	                f[j+1] = max(f[j+1], f[j])
	            }
	        }
	    }
	    return f[m]
	}

时间复杂度: o(nm)
空间复杂度: o(m)
*/
func main() {

}
