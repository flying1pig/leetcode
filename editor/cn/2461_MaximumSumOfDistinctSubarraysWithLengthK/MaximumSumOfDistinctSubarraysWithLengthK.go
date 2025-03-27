package main

/**
给你一个整数数组 nums 和一个整数 k 。请你从 nums 中满足下述条件的全部子数组中找出最大子数组和：


 子数组的长度是 k，且
 子数组中的所有元素 各不相同 。


 返回满足题面要求的最大子数组和。如果不存在子数组满足这些条件，返回 0 。

 子数组 是数组中一段连续非空的元素序列。



 示例 1：

 输入：nums = [1,5,4,2,9,9,9], k = 3
输出：15
解释：nums 中长度为 3 的子数组是：
- [1,5,4] 满足全部条件，和为 10 。
- [5,4,2] 满足全部条件，和为 11 。
- [4,2,9] 满足全部条件，和为 15 。
- [2,9,9] 不满足全部条件，因为元素 9 出现重复。
- [9,9,9] 不满足全部条件，因为元素 9 出现重复。
因为 15 是满足全部条件的所有子数组中的最大子数组和，所以返回 15 。


 示例 2：

 输入：nums = [4,4,4], k = 3
输出：0
解释：nums 中长度为 3 的子数组是：
- [4,4,4] 不满足全部条件，因为元素 4 出现重复。
因为不存在满足全部条件的子数组，所以返回 0 。




 提示：


 1 <= k <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁵


 Related Topics 数组 哈希表 滑动窗口 👍 74 👎 0

*/

/*
题型: 定长滑动窗口
题目: 长度为 K 子数组中的最大和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSubarraySum(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}

//时间复杂度: o(n)
//空间复杂度: o(k)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
