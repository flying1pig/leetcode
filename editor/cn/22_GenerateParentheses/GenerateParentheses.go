package main

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



 示例 1：


输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]


 示例 2：


输入：n = 1
输出：["()"]




 提示：


 1 <= n <= 8


 Related Topics 字符串 动态规划 回溯 👍 3807 👎 0

*/

/*
题型: 回溯
题目: 括号生成
*/

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	// i = 目前填了多少个括号
	// open = 左括号个数，i-open = 右括号个数
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
