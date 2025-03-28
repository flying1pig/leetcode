package main

/**
给定两个整数 a 和 b ，返回 任意 字符串 s ，要求满足：


 s 的长度为 a + b，且正好包含 a 个 'a' 字母与 b 个 'b' 字母；
 子串 'aaa' 没有出现在 s 中；
 子串 'bbb' 没有出现在 s 中。




 示例 1：


输入：a = 1, b = 2
输出："abb"
解释："abb", "bab" 和 "bba" 都是正确答案。


 示例 2：


输入：a = 4, b = 1
输出："aabaa"



 提示：


 0 <= a, b <= 100
 对于给定的 a 和 b，保证存在满足要求的 s



 Related Topics 贪心 字符串 👍 101 👎 0

*/

/*
题型: 贪心
题目: 吃苹果的最大数目
*/

// leetcode submit region begin(Prohibit modification and deletion)
func strWithout3a3b(a int, b int) string {
	ans := ""
	for a != 0 && b != 0 {
		if a > b {
			ans += "aab"
			a = a - 2
			b--
		} else if b > a {
			ans += "bba"
			a--
			b = b - 2
		} else {
			ans += "ab"
			a--
			b--
		}
	}
	if a == 0 {
		for b != 0 {
			ans += "b"
			b--
		}
	}
	if b == 0 {
		for a != 0 {
			ans += "a"
			a--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
