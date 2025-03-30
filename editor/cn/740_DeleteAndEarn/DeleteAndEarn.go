package main

import "slices"

/**
给你一个整数数组 nums ，你可以对它进行一些操作。

 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] +
 1 的元素。

 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。



 示例 1：


输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。


 示例 2：


输入：nums = [2,2,3,3,3,4]
输出：9
解释：
删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。




 提示：


 1 <= nums.length <= 2 * 10⁴
 1 <= nums[i] <= 10⁴


 👍 1088 👎 0

*/

/*
题型: dp
题目: 删除并获得点数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func deleteAndEarn(nums []int) int {
	rob := func(nums []int) int {
		m0 := 0
		m1 := nums[0]
		for i := 2; i <= len(nums); i++ {
			newm := max(m1, m0+nums[i-1])
			m0 = m1
			m1 = newm
		}
		return m1
	}

	a := make([]int, slices.Max(nums)+1)
	for _, val := range nums {
		a[val] += val
	}
	return rob(a)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路: 把nums数组转换为值域数组a，其中a[i]表示nums中的等于i的数组元素之和。剩下的做法同198_HouseRobber。
func deleteAndEarn(nums []int) int {
    a := make([]int, slices.Max(nums)+1)
    for _,val := range nums {
        a[val] += val
    }
    return rob(a)
}

func rob(nums []int) int {
    m0 := 0
    m1 := nums[0]
    for i := 2; i <= len(nums); i++ {
        newm := max(m1, m0+nums[i-1])
        m0 = m1
        m1 = newm
    }
    return m1
}
时间复杂度: o(n+u), n是数组长度，u=max(nums)
空间复杂度: o(u)
*/

func main() {

}
