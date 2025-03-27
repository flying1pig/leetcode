package main

import "slices"

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

 注意：答案中不可以包含重复的三元组。





 示例 1：


输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。


 示例 2：


输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。


 示例 3：


输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。




 提示：


 3 <= nums.length <= 3000
 -10⁵ <= nums[i] <= 10⁵


 Related Topics 数组 双指针 排序 👍 7395 👎 0

*/

/*
题型: 相向双指针
题目: 三数之和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // 跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化一
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // 优化二
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else { // 三数之和为 0
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}
		}
	}
	return
}

//时间复杂度：O(n^2)，其中 n 为 nums 的长度
//空间复杂度：O(1)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
