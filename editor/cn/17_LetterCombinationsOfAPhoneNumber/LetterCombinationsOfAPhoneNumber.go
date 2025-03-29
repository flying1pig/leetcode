package main

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





 示例 1：


输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]


 示例 2：


输入：digits = ""
输出：[]


 示例 3：


输入：digits = "2"
输出：["a","b","c"]




 提示：


 0 <= digits.length <= 4
 digits[i] 是范围 ['2', '9'] 的一个数字。


 Related Topics 哈希表 字符串 回溯 👍 3032 👎 0

*/

/*
题型: 回溯
题目: 电话号码的字母组合
*/

// leetcode submit region begin(Prohibit modification and deletion)
var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	path := make([]byte, n) // 注意 path 长度一开始就是 n，不是空列表
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			path[i] = byte(c) // 直接覆盖
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
