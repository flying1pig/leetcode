package main

/**
给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。

 你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：


 如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
 对于你访问的位置 i ，你可以获得分数 nums[i] 。
 如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。


 请你返回你能得到的 最大 得分之和。

 注意 ，你一开始的分数为 nums[0] 。



 示例 1：

 输入：nums = [2,3,6,1,9,2], x = 5
输出：13
解释：我们可以按顺序访问数组中的位置：0 -> 2 -> 3 -> 4 。
对应位置的值为 2 ，6 ，1 和 9 。因为 6 和 1 的奇偶性不同，所以下标从 2 -> 3 让你失去 x = 5 分。
总得分为：2 + 6 + 1 + 9 - 5 = 13 。


 示例 2：

 输入：nums = [2,4,6,8], x = 3
输出：20
解释：数组中的所有元素奇偶性都一样，所以我们可以将每个元素都访问一次，而且不会失去任何分数。
总得分为：2 + 4 + 6 + 8 = 20 。




 提示：


 2 <= nums.length <= 10⁵
 1 <= nums[i], x <= 10⁶


 Related Topics 数组 动态规划 👍 63 👎 0

*/

/*
题型: dp
题目: 访问数组中的位置使分数最大
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(nums []int, x int) int64 {
	n := len(nums)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if nums[i]%2 != j {
			return dfs(i+1, j)
		}
		return max(dfs(i+1, j), dfs(i+1, j^1)-x) + nums[i]
	}
	return int64(dfs(0, nums[0]%2))
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	dfs(i,j)=max(dfs(i+1,j),dfs(i+1,j⊕1)−x)+v
*/

/*
递推

	func maxScore(nums []int, x int) int64 {
	    n := len(nums)
	    f := make([][2]int, n+1)
	    for i := n - 1; i >= 0; i-- {
	        v := nums[i]
	        r := v % 2
	        f[i][r^1] = f[i+1][r^1] // v%2 != j 的情况
	        f[i][r] = max(f[i+1][r], f[i+1][r^1]-x) + v
	    }
	    return int64(f[0][nums[0]%2])
	}
*/

/*
空间优化

	func maxScore(nums []int, x int) int64 {
	    f := [2]int{}
	    for i := len(nums) - 1; i >= 0; i-- {
	        v := nums[i]
	        r := v % 2
	        f[r] = max(f[r], f[r^1]-x) + v
	    }
	    return int64(f[nums[0]%2])
	}
*/
func main() {

}
