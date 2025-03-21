package main

import "math"

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

 子数组 是数组中的一个连续部分。



 示例 1：


输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。


 示例 2：


输入：nums = [1]
输出：1


 示例 3：


输入：nums = [5,4,-1,7,8]
输出：23




 提示：


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴




 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

 👍 6976 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思考:
定义f(i)为以nums[i]结尾的子数组的最大和。
当nums[i]选，f(i) = f(i-1)+nums[i]
当nums[i]不选，f(i) = f(i-1)
所以，f(i) = max(f(i-1),f(i-1)+nums[i]) = max(0,f(i-1)) + nums[i]
是否要判断子数组是否连续？不需要。

状态方程:
f(i) = max(0,f(i-1)) + nums[i]
边界条件:
f(0) = nums[0]

时间复杂度: o(n)
空间复杂度: o(1)
*/
func main() {

}
