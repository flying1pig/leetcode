package main

import "slices"

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋
装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。



 示例 1：


输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。


 示例 2：


输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

 示例 3：


输入：nums = [1,2,3]
输出：3




 提示：


 1 <= nums.length <= 100
 0 <= nums[i] <= 1000


 👍 1681 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if len(nums) <= 3 {
		return slices.Max(nums)
	}
	f := func(arr []int) int {
		m0 := 0
		m1 := arr[0]
		for i := 2; i <= len(arr); i++ {
			newm := max(m1, m0+arr[i-1])
			m0 = m1
			m1 = newm
		}
		return m1
	}
	return max(f(nums[2:len(nums)-1])+nums[0], f(nums[1:len(nums)]))
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
    下标为0的房子偷，子问题变为: f(i) = max(f(i-1),f(i-2)+nums[i]) + nums[0], 2<=i<=n-2。此时问题简化为198
    下标为0的房子不偷, 子问题变为: f(i) = max(f(i-1),f(i-2)+nums[i]), 1<=i<=n-1。此时问题简化为198
    答案为取上述两种情况的最大者
边界条件:

*/

func main() {

}
