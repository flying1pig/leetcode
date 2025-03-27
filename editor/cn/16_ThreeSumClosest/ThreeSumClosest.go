package main

import (
	"math"
	"sort"
)

/**
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

 返回这三个数的和。

 假定每组输入只存在恰好一个解。



 示例 1：


输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。


 示例 2：


输入：nums = [0,0,0], target = 1
输出：0
解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。



 提示：


 3 <= nums.length <= 1000
 -1000 <= nums[i] <= 1000
 -10⁴ <= target <= 10⁴


 Related Topics 数组 双指针 排序 👍 1715 👎 0

*/

/*
题型: 相向双指针
题目: 最接近的三数之和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	minDiff := math.MaxInt
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue // 优化三
		}

		// 优化一
		s := x + nums[i+1] + nums[i+2]
		if s > target { // 后面无论怎么选，选出的三个数的和不会比 s 还小
			if s-target < minDiff {
				ans = s // 由于下面直接 break，这里无需更新 minDiff
			}
			break
		}

		// 优化二
		s = x + nums[n-2] + nums[n-1]
		if s < target { // x 加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了
			if target-s < minDiff {
				minDiff = target - s
				ans = s
			}
			continue
		}

		// 双指针
		j, k := i+1, n-1
		for j < k {
			s = x + nums[j] + nums[k]
			if s == target {
				return target
			}
			if s > target {
				if s-target < minDiff { // s 与 target 更近
					minDiff = s - target
					ans = s
				}
				k--
			} else { // s < target
				if target-s < minDiff { // s 与 target 更近
					minDiff = target - s
					ans = s
				}
				j++
			}
		}
	}
	return ans
}

//时间复杂度: o(n^2)
//空间复杂度: o(1)

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
