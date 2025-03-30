package main

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列
。

 示例 1：


输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。


 示例 2：


输入：nums = [0,1,0,3,2,3]
输出：4


 示例 3：


输入：nums = [7,7,7,7,7,7,7]
输出：1




 提示：


 1 <= nums.length <= 2500
 -10⁴ <= nums[i] <= 10⁴




 进阶：


 你能将算法的时间复杂度降低到 O(n log(n)) 吗?


 Related Topics 数组 二分查找 动态规划 👍 3929 👎 0

*/

/*
题型: dp
题目: 最长递增子序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n) // 本题可以初始化成 0，表示没有计算过
	var dfs func(int) int
	dfs = func(i int) int {
		p := &memo[i]
		if *p > 0 {
			return *p
		}
		res := 0
		for j, x := range nums[:i] {
			if x < nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++ // 加一提到循环外面
		*p = res
		return res
	}
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs(i)表示以nums[i]为结尾的最长递增子序列的长度
枚举子序列倒数第二个数的下标j, 如果nums[j] < nums[i]，那么有dfs(i) = dfs(j)+1
*/

/*
记忆化搜索:
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n) // 本题可以初始化成 0，表示没有计算过
	var dfs func(int) int
	dfs = func(i int) int {
		p := &memo[i]
		if *p > 0 {
			return *p
		}
		res := 0
		for j, x := range nums[:i] {
			if x < nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++ // 加一提到循环外面
		*p = res
		return res
	}
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return
}
时间复杂度: o(n^2)
空间复杂度: o(n)
*/

/*
递推:

	func lengthOfLIS(nums []int) int {
	    f := make([]int, len(nums))
	    for i, x := range nums {
	        for j, y := range nums[:i] {
	            if y < x {
	                f[i] = max(f[i], f[j])
	            }
	        }
	        f[i]++
	    }
	    return slices.Max(f)
	}

时间复杂度: o(n^2)
空间复杂度: o(n)
*/
func main() {

}
