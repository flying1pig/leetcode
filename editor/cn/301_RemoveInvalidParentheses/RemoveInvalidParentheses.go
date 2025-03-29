package main

/**
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

 返回所有可能的结果。答案可以按 任意顺序 返回。



 示例 1：


输入：s = "()())()"
输出：["(())()","()()()"]


 示例 2：


输入：s = "(a)())()"
输出：["(a())()","(a)()()"]


 示例 3：


输入：s = ")("
输出：[""]




 提示：


 1 <= s.length <= 25
 s 由小写英文字母以及括号 '(' 和 ')' 组成
 s 中至多含 20 个括号


 Related Topics 广度优先搜索 字符串 回溯 👍 969 👎 0

*/

/*
题型: 回溯
题目: 删除无效的括号
*/

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func removeInvalidParentheses(s string) (ans []string) {
	curSet := map[string]struct{}{s: {}}
	for {
		for str := range curSet {
			if isValid(str) {
				ans = append(ans, str)
			}
		}
		if len(ans) > 0 {
			return
		}
		nextSet := map[string]struct{}{}
		for str := range curSet {
			for i, ch := range str {
				if i > 0 && byte(ch) == str[i-1] {
					continue
				}
				if ch == '(' || ch == ')' {
					nextSet[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		curSet = nextSet
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
