package main

/**
给你一个二进制数组 nums ，你需要从中删掉一个元素。

 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

 如果不存在这样的子数组，请返回 0 。



 提示 1：


输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。

 示例 2：


输入：nums = [0,1,1,1,0,1,1,0,1]
输出：5
解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。

 示例 3：


输入：nums = [1,1,1]
输出：2
解释：你必须要删除一个元素。



 提示：


 1 <= nums.length <= 10⁵
 nums[i] 要么是 0 要么是 1 。


 Related Topics 数组 动态规划 滑动窗口 👍 176 👎 0

*/

/*
题型: 不定长滑动窗口
题目: 删掉一个元素以后全为 1 的最长子数组
*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestSubarray(nums []int) int {
	//窗口中可以包含一个0
	idx := -1
	ans := 0
	l := 0
	for r, v := range nums {
		//首次遇到
		if v == 0 && idx == -1 {
			idx = r
			continue
		}
		//遇到第二个0
		if v == 0 {
			ans = max(ans, r-l)
			//更新到第一个0后的第一个1的下标
			l = idx + 1
			//更新到当前0
			idx = r
		}

	}
	ans = max(ans, len(nums)-l) - 1
	return ans
}

//时间复杂度: o(n)
//空间复杂度: o(1)

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
