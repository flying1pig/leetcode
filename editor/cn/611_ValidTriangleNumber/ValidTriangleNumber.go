package main

import "slices"

/**
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。



 示例 1:


输入: nums = [2,2,3,4]
输出: 3
解释:有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3


 示例 2:


输入: nums = [4,2,3,4]
输出: 4



 提示:


 1 <= nums.length <= 1000
 0 <= nums[i] <= 1000


 Related Topics 贪心 数组 双指针 二分查找 排序 👍 617 👎 0

*/
/*
题型: 相向双指针
题目: 有效三角形的个数
*/
//leetcode submit region begin(Prohibit modification and deletion)
func triangleNumber(nums []int) (ans int) {
	slices.Sort(nums)
	for k := 2; k < len(nums); k++ {
		c := nums[k]
		i := 0     // a=nums[i]
		j := k - 1 // b=nums[j]
		for i < j {
			if nums[i]+nums[j] > c {
				// 由于 nums 已经从小到大排序
				// nums[i]+nums[j] > c 同时意味着：
				// nums[i+1]+nums[j] > c
				// nums[i+2]+nums[j] > c
				// ...
				// nums[j-1]+nums[j] > c
				// 从 i 到 j-1 一共 j-i 个
				ans += j - i
				j--
			} else {
				// 由于 nums 已经从小到大排序
				// nums[i]+nums[j] <= c 同时意味着
				// nums[i]+nums[j-1] <= c
				// ...
				// nums[i]+nums[i+1] <= c
				// 所以在后续的内层循环中，nums[i] 不可能作为三角形的边长，没有用了
				i++
			}
		}
	}
	return
}

//时间复杂度: o(n^2)
//空间复杂度: o(1)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
