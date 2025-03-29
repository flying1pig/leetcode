package main

import (
	"math"
	"strconv"
)

/**
给定一个数字字符串 num，比如 "123456579"，我们可以将它分成「斐波那契式」的序列 [123, 456, 579]。

 形式上，斐波那契式 序列是一个非负整数列表 f，且满足：


 0 <= f[i] < 2³¹ ，（也就是说，每个整数都符合 32 位 有符号整数类型）
 f.length >= 3
 对于所有的0 <= i < f.length - 2，都有 f[i] + f[i + 1] = f[i + 2]


 另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。

 返回从 num 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。



 示例 1：


输入：num = "1101111"
输出：[11,0,11,11]
解释：输出 [110,1,111] 也可以。

 示例 2：


输入: num = "112358130"
输出: []
解释: 无法拆分。


 示例 3：


输入："0123"
输出：[]
解释：每个块的数字不能以零开头，因此 "01"，"2"，"3" 不是有效答案。




 提示：


 1 <= num.length <= 200
 num 中只含有数字


 Related Topics 字符串 回溯 👍 304 👎 0

*/

/*
题型: 回溯
题目: 将数组拆分成斐波那契序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
// dfs
func splitIntoFibonacci(S string) []int {

	var res []int
	var dfs func(idx int, nums []int)
	dfs = func(idx int, nums []int) {
		if len(res) > 2 {
			return
		}

		if idx == len(S) && len(nums) > 2 {
			res = make([]int, len(nums))
			copy(res, nums)
			return
		}

		for i := idx; i < len(S); i++ {
			cur, err := strconv.Atoi(S[idx : i+1])
			if err != nil || cur > math.MaxInt32 {
				break
			}
			if len(nums) < 2 {
				nums = append(nums, cur)
				dfs(i+1, nums)
				nums = nums[:len(nums)-1]
			} else if cur == nums[len(nums)-1]+nums[len(nums)-2] {
				nums = append(nums, cur)
				dfs(i+1, nums)
				nums = nums[:len(nums)-1]
			}
			if S[idx] == '0' {
				break
			}
		}
	}
	dfs(0, []int{})

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
