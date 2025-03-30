package main

import "slices"

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

 测试用例的答案是一个 32-位 整数。



 示例 1:


输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。


 示例 2:


输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。



 提示:


 1 <= nums.length <= 2 * 10⁴
 -10 <= nums[i] <= 10
 nums 的任何子数组的乘积都 保证 是一个 32-位 整数


 Related Topics 数组 动态规划 👍 2407 👎 0

*/

/*
题型: dp
题目: 乘积最大子数组
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	n := len(nums)
	fMax := make([]int, n)
	fMin := make([]int, n)
	fMax[0], fMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		// 把 x 加到右端点为 i-1 的（乘积最大/最小）子数组后面，
		// 或者单独组成一个子数组，只有 x 一个元素
		fMax[i] = max(fMax[i-1]*x, fMin[i-1]*x, x)
		fMin[i] = min(fMax[i-1]*x, fMin[i-1]*x, x)
	}
	return slices.Max(fMax)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
fmax[i]=max(fmax[i−1]⋅x,fmin[i−1]⋅x,x)
fmin[i]=min(fmax[i−1]⋅x,fmin[i−1]⋅x,x)
*/

func main() {

}
