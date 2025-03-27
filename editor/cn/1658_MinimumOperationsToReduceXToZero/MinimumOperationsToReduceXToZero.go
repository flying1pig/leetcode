package main

/**
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数
组以供接下来的操作使用。

 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。



 示例 1：


输入：nums = [1,1,4,2,3], x = 5
输出：2
解释：最佳解决方案是移除后两个元素，将 x 减到 0 。


 示例 2：


输入：nums = [5,6,7,8,9], x = 4
输出：-1


 示例 3：


输入：nums = [3,2,20,1,1,3], x = 10
输出：5
解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。




 提示：


 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁴
 1 <= x <= 10⁹


 Related Topics 数组 哈希表 二分查找 前缀和 滑动窗口 👍 411 👎 0

*/
/*
题型: 不定长滑动窗口
题目: 将 x 减到 0 的最小操作数
*/
//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, x int) int {
	target := -x
	for _, x := range nums {
		target += x
	}
	if target < 0 { // 全部移除也无法满足要求
		return -1
	}

	ans, left, sum := -1, 0, 0
	for right, x := range nums {
		sum += x
		for sum > target {
			sum -= nums[left]
			left++ // 缩小子数组长度
		}
		if sum == target {
			ans = max(ans, right-left+1)
		}
	}
	if ans < 0 {
		return -1
	}
	return len(nums) - ans
}

//时间复杂度: o(n)
//空间复杂度: o(1)
//leetcode submit region end(Prohibit modification and deletion)

/*
提示: 逆向思维+滑动窗口
把问题转换成「从 nums 中移除一个最长的子数组，使得剩余元素的和为 x」
*/

func main() {

}
