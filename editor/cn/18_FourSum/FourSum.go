package main

import "slices"

/**
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b],
 nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：


 0 <= a, b, c, d < n
 a、b、c 和 d 互不相同
 nums[a] + nums[b] + nums[c] + nums[d] == target


 你可以按 任意顺序 返回答案 。



 示例 1：


输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]


 示例 2：


输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]




 提示：


 1 <= nums.length <= 200
 -10⁹ <= nums[i] <= 10⁹
 -10⁹ <= target <= 10⁹


 Related Topics 数组 双指针 排序 👍 2043 👎 0

*/

/*
题型: 相向双指针
题目: 四数之和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for a := range n - 3 { // 枚举第一个数
		x := nums[a]
		if a > 0 && x == nums[a-1] { // 跳过重复数字
			continue
		}
		if x+nums[a+1]+nums[a+2]+nums[a+3] > target { // 优化一
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target { // 优化二
			continue
		}
		for b := a + 1; b < n-2; b++ { // 枚举第二个数
			y := nums[b]
			if b > a+1 && y == nums[b-1] { // 跳过重复数字
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target { // 优化一
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target { // 优化二
				continue
			}
			c, d := b+1, n-1
			for c < d { // 双指针枚举第三个数和第四个数
				s := x + y + nums[c] + nums[d] // 四数之和
				if s > target {
					d--
				} else if s < target {
					c++
				} else { // s == target
					ans = append(ans, []int{x, y, nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					} // 跳过重复数字
					for d--; d > c && nums[d] == nums[d+1]; d-- {
					} // 跳过重复数字
				}
			}
		}
	}
	return
}

//时间复杂度：O(n^3)，其中 n 为 nums 的长度。排序 O(nlogn)。两重循环枚举第一个数和第二个数，然后 O(n) 双指针枚举第三个数和第四个数。所以总的时间复杂度为 O(n^3)。
//空间复杂度：O(1)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
