package main

import "strconv"

/**
当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。

 给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。



 示例 1:


输入: n = 10
输出: 9


 示例 2:


输入: n = 1234
输出: 1234


 示例 3:


输入: n = 332
输出: 299




 提示:


 0 <= n <= 10⁹


 Related Topics 贪心 数学 👍 511 👎 0

*/

/*
题型: 字符串贪心
题目: 单调递增的数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
func monotoneIncreasingDigits(n int) int {
	strN := strconv.Itoa(n)
	ss := []byte(strN)
	size := len(ss)
	if size <= 1 {
		return n
	}
	for i := size - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1]--
			for j := i; j < size; j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
