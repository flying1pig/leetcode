package main

/**
给你一个正整数数组 nums ，请你从中删除一个含有 若干不同元素 的子数组。删除子数组的 得分 就是子数组各元素之 和 。

 返回 只删除一个 子数组可获得的 最大得分 。

 如果数组 b 是数组 a 的一个连续子序列，即如果它等于 a[l],a[l+1],...,a[r] ，那么它就是 a 的一个子数组。



 示例 1：


输入：nums = [4,2,4,5,6]
输出：17
解释：最优子数组是 [2,4,5,6]


 示例 2：


输入：nums = [5,2,1,2,5,2,1,2,5]
输出：8
解释：最优子数组是 [5,2,1] 或 [1,2,5]




 提示：


 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁴


 Related Topics 数组 哈希表 滑动窗口 👍 107 👎 0

*/

/*
题型: 不定长滑动窗口
题目: 删除子数组的最大得分
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumUniqueSubarray(nums []int) (maxSum int) {
	set := map[int]bool{}
	for start, end, sum := 0, 0, 0; end < len(nums); end++ {
		for set[nums[end]] { // 因重复而剔除start元素
			sum -= nums[start]
			delete(set, nums[start])
			start++
		}
		sum += nums[end]          // 加入end元素后的sum
		maxSum = max(maxSum, sum) // 求max
		set[nums[end]] = true     // 加入end元素
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
