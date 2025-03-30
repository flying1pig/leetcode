package main

/**
给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。

 每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。

 请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。



 示例 1：

 输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
输出：3
解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
- 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
- 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
- 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。


 示例 2：

 输入：nums1 = [1,1,1,1,1,1,1], nums2 = [6]
输出：-1
解释：没有办法减少 nums1 的和或者增加 nums2 的和使二者相等。


 示例 3：

 输入：nums1 = [6,6], nums2 = [1]
输出：3
解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
- 将 nums1[0] 变为 2 。 nums1 = [2,6], nums2 = [1] 。
- 将 nums1[1] 变为 2 。 nums1 = [2,2], nums2 = [1] 。
- 将 nums2[0] 变为 4 。 nums1 = [2,2], nums2 = [4] 。




 提示：


 1 <= nums1.length, nums2.length <= 10⁵
 1 <= nums1[i], nums2[i] <= 6


 Related Topics 贪心 数组 哈希表 计数 👍 200 👎 0

*/

/*
题型: 贪心
题目: 通过最少操作次数使数组的和相等
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums1, nums2 []int) (ans int) {
	if 6*len(nums1) < len(nums2) || 6*len(nums2) < len(nums1) {
		return -1 // 优化
	}
	d := 0 // 数组元素和的差，我们要让这个差变为 0
	for _, x := range nums2 {
		d += x
	}
	for _, x := range nums1 {
		d -= x
	}
	if d < 0 {
		d = -d
		nums1, nums2 = nums2, nums1 // 统一让 nums1 的数变大，nums2 的数变小
	}
	cnt := [6]int{} // 统计每个数的最大变化量
	for _, x := range nums1 {
		cnt[6-x]++
	} // nums1 的变成 6
	for _, x := range nums2 {
		cnt[x-1]++
	} // nums2 的变成 1
	for i := 5; ; i-- { // 从大到小枚举最大变化量 5 4 3 2 1
		if i*cnt[i] >= d { // 可以让 d 变为 0
			return ans + (d+i-1)/i
		}
		ans += cnt[i] // 需要所有最大变化量为 i 的数
		d -= i * cnt[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
