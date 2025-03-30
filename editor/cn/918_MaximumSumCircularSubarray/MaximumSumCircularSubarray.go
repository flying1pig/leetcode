package main

import "math"

/**
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素
是 nums[(i - 1 + n) % n] 。

 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在
 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。



 示例 1：


输入：nums = [1,-2,3,-2]
输出：3
解释：从子数组 [3] 得到最大和 3


 示例 2：


输入：nums = [5,-3,5]
输出：10
解释：从子数组 [5,5] 得到最大和 5 + 5 = 10


 示例 3：


输入：nums = [3,-2,2,-3]
输出：3
解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3




 提示：


 n == nums.length
 1 <= n <= 3 * 10⁴
 -3 * 10⁴ <= nums[i] <= 3 * 10⁴


 👍 792 👎 0

*/

/*
题型: dp
题目: 环形子数组的最大和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(nums []int) int {
	maxS := math.MinInt //最大子数组和，不能为空
	minS := 0           //最小子数组和，可以为空
	maxF, minF, sum := 0, 0, 0
	for _, x := range nums {
		maxF = max(maxF, 0) + x
		maxS = max(maxS, maxF)
		minF = min(minF, 0) + x
		minS = min(minS, minF)
		sum += x
	}
	if sum == minS {
		return maxS
	}
	return max(maxS, sum-minS)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
先看提示
case1：最大连续子数组在数组中间
case2：最大连续子数组包含环形连接处，那么sum(nums) - 最小连续子数组就得到最大
*/

func main() {

}
