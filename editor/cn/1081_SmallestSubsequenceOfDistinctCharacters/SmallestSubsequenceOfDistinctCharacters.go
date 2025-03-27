package main

/**
返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。



 示例 1：


输入：s = "bcabc"
输出："abc"


 示例 2：


输入：s = "cbacdcbc"
输出："acdb"



 提示：


 1 <= s.length <= 1000
 s 由小写英文字母组成




 注意：该题与 316 https://leetcode.cn/problems/remove-duplicate-letters/ 相同

 Related Topics 栈 贪心 字符串 单调栈 👍 224 👎 0

*/

/*
题型: 单调栈
题目: 不同字符的最小子序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
func smallestSubsequence(s string) string {
	left := ['z' + 1]int{} // 相比创建一个长为 26 的数组，多开一点空间更方便
	for _, c := range s {
		left[c]++ // 统计每个字母的出现次数
	}
	ans := []rune{}
	inAns := ['z' + 1]bool{}
	for _, c := range s {
		left[c]--
		if inAns[c] { // ans 中不能有重复字母
			continue
		}
		for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
			// 如果 c < x，且右边还有 x，那么可以把 x 去掉，因为后面可以重新把 x 加到 ans 中
			x := ans[len(ans)-1]
			ans = ans[:len(ans)-1]
			inAns[x] = false // 标记 x 不在 ans 中
		}
		ans = append(ans, c) // 把 c 加到 ans 的末尾
		inAns[c] = true      // 标记 c 在 ans 中
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
