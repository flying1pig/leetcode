package main

/**
给你一个字符串 s，最多 可以从中删除一个字符。

 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。



 示例 1：


输入：s = "aba"
输出：true


 示例 2：


输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。


 示例 3：


输入：s = "abc"
输出：false



 提示：


 1 <= s.length <= 10⁵
 s 由小写英文字母组成


 Related Topics 贪心 双指针 字符串 👍 698 👎 0

*/

/*
题型: 字符串贪心
题目: 验证回文串 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			// 删除 s[i] 或者 s[j]（注意 Go 的切片是 O(1) 的，不会生成新字符串）
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
		i++
		j--
	}
	return true // s 本身就是回文串
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
