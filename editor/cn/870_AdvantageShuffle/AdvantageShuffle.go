package main

import "slices"

/**
给定两个长度相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目
来描述。

 返回 nums1 的 任意 排列，使其相对于 nums2 的优势最大化。



 示例 1：


输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
输出：[2,11,7,15]


 示例 2：


输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
输出：[24,32,8,12]




 提示：


 1 <= nums1.length <= 10⁵
 nums2.length == nums1.length
 0 <= nums1[i], nums2[i] <= 10⁹


 Related Topics 贪心 数组 双指针 排序 👍 451 👎 0

*/

/*
题型: 贪心
题目: 优势洗牌
*/

// leetcode submit region begin(Prohibit modification and deletion)
func advantageCount(nums1, nums2 []int) []int {
	slices.Sort(nums1)

	n := len(nums1)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums2[i] - nums2[j] })

	ans := make([]int, n)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[idx[left]] {
			ans[idx[left]] = x // 用下等马比下等马
			left++
		} else {
			ans[idx[right]] = x // 用下等马比上等马
			right--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
