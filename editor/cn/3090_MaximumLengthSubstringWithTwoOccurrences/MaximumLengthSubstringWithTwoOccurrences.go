package main

/**
给你一个字符串 s ，请找出满足每个字符最多出现两次的最长子字符串，并返回该子字符串的 最大 长度。



 示例 1：


 输入： s = "bcbbbcba"


 输出： 4

 解释：

 以下子字符串长度为 4，并且每个字符最多出现两次："bcbbbcba"。

 示例 2：


 输入： s = "aaaa"


 输出： 2

 解释：

 以下子字符串长度为 2，并且每个字符最多出现两次："aaaa"。



 提示：



 2 <= s.length <= 100

 s 仅由小写英文字母组成。


 Related Topics 哈希表 字符串 滑动窗口 👍 22 👎 0

*/

/*
题型: 不定长滑动窗口
题目: 每个字符最多出现两次的最长子字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLengthSubstring(s string) (ans int) {
	cnt := [26]int{}
	left := 0
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		for cnt[b] > 2 {
			cnt[s[left]-'a']--
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}

//时间复杂度:	o(n+x), n为s长度，x为字符集合大小，本题均为小写字符，所以x=26
//空间复杂度: o(x)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
