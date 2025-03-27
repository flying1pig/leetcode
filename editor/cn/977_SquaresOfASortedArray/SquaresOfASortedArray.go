package main

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。






 示例 1：


输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

 示例 2：


输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]




 提示：


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums 已按 非递减顺序 排序




 进阶：


 请你设计时间复杂度为 O(n) 的算法解决本问题


 Related Topics 数组 双指针 排序 👍 1093 👎 0

*/

/*
题型: 相向双指针
题目: 有序数组的平方
*/

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for p := n - 1; p >= 0; p-- {
		x := nums[i] * nums[i]
		y := nums[j] * nums[j]
		if x > y {
			ans[p] = x
			i++
		} else {
			ans[p] = y
			j--
		}
	}
	return ans
}

//时间复杂度: o(n)
//空间复杂度: o(1)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
