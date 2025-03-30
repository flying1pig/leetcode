package main

import (
	"sort"
)

/**
给你一个整数数组 nums，请你找出并返回能被三整除的元素 最大和。






 示例 1：


输入：nums = [3,6,5,1,8]
输出：18
解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。

 示例 2：


输入：nums = [4]
输出：0
解释：4 不能被 3 整除，所以无法选出数字，返回 0。


 示例 3：


输入：nums = [1,2,3,4,4]
输出：12
解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。




 提示：


 1 <= nums.length <= 4 * 10⁴
 1 <= nums[i] <= 10⁴


 Related Topics 贪心 数组 动态规划 排序 👍 367 👎 0

*/

/*
题型: 贪心
题目: 可被三整除的最大和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumDivThree(nums []int) (ans int) {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%3 == 0 {
		return s
	}

	a := [3][]int{}
	for _, x := range nums {
		a[x%3] = append(a[x%3], x)
	}
	sort.Ints(a[1])
	sort.Ints(a[2])

	if s%3 == 2 {
		a[1], a[2] = a[2], a[1]
	}
	if len(a[1]) > 0 {
		ans = s - a[1][0]
	}
	if len(a[2]) > 1 {
		ans = max(ans, s-a[2][0]-a[2][1])
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
