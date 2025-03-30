package main

import (
	"math/bits"
	"slices"
)

/**
给你一个长度为 3 的整数数组 nums。

 现以某种顺序 连接 数组 nums 中所有元素的 二进制表示 ，请你返回可以由这种方法形成的 最大 数值。

 注意 任何数字的二进制表示 不含 前导零。



 示例 1:


 输入: nums = [1,2,3]


 输出: 30

 解释:

 按照顺序 [3, 1, 2] 连接数字的二进制表示，得到结果 "11110"，这是 30 的二进制表示。

 示例 2:


 输入: nums = [2,8,16]


 输出: 1296

 解释:

 按照顺序 [2, 8, 16] 连接数字的二进制表述，得到结果 "10100010000"，这是 1296 的二进制表示。



 提示:


 nums.length == 3
 1 <= nums[i] <= 127


 Related Topics 位运算 数组 枚举 👍 4 👎 0

*/

/*
题型: 贪心
题目: 连接二进制表示可形成的最大数值
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxGoodNumber(nums []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int {
		lenA := bits.Len(uint(a))
		lenB := bits.Len(uint(b))
		return (b<<lenA | a) - (a<<lenB | b)
	})

	for _, x := range nums {
		ans = ans<<bits.Len(uint(x)) | x
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
