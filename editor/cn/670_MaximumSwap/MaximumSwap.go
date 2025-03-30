package main

import "strconv"

/**
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

 示例 1 :


输入: 2736
输出: 7236
解释: 交换数字2和数字7。


 示例 2 :


输入: 9973
输出: 9973
解释: 不需要交换。


 注意:


 给定数字的范围是 [0, 10⁸]


 Related Topics 贪心 数学 👍 516 👎 0

*/

/*
题型: 字符串贪心
题目: 最大交换
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	maxIdx := len(s) - 1
	p, q := -1, 0
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[maxIdx] { // s[i] 是目前最大数字
			maxIdx = i
		} else if s[i] < s[maxIdx] { // s[i] 右边有比它大的
			p, q = i, maxIdx // 更新 p 和 q
		}
	}
	if p == -1 { // 这意味着 s 是降序的
		return num
	}
	t := []byte(s)
	t[p], t[q] = t[q], t[p]
	ans, _ := strconv.Atoi(string(t))
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
