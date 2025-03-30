package main

import "math"

/**
给你一个长度为 n 的整数数组 nums。

 子数组 nums[l..r]（其中 0 <= l <= r < n）的 成本 定义为：

 cost(l, r) = nums[l] - nums[l + 1] + ... + nums[r] * (−1)r − l

 你的任务是将 nums 分割成若干子数组，使得所有子数组的成本之和 最大化，并确保每个元素 正好 属于一个子数组。

 具体来说，如果 nums 被分割成 k 个子数组，且分割点为索引 i1, i2, ..., ik − 1（其中 0 <= i1 < i2 < ... <
ik - 1 < n - 1），则总成本为：

 cost(0, i1) + cost(i1 + 1, i2) + ... + cost(ik − 1 + 1, n − 1)

 返回在最优分割方式下的子数组成本之和的最大值。

 注意：如果 nums 没有被分割，即 k = 1，则总成本即为 cost(0, n - 1)。



 示例 1：


 输入： nums = [1,-2,3,4]


 输出： 10

 解释：

 一种总成本最大化的方法是将 [1, -2, 3, 4] 分割成子数组 [1, -2, 3] 和 [4]。总成本为 (1 + 2 + 3) + 4 = 10。


 示例 2：


 输入： nums = [1,-1,1,-1]


 输出： 4

 解释：

 一种总成本最大化的方法是将 [1, -1, 1, -1] 分割成子数组 [1, -1] 和 [1, -1]。总成本为 (1 + 1) + (1 + 1) =
4。

 示例 3：


 输入： nums = [0]


 输出： 0

 解释：

 无法进一步分割数组，因此答案为 0。

 示例 4：


 输入： nums = [1,-1]


 输出： 2

 解释：

 选择整个数组，总成本为 1 + 1 = 2，这是可能的最大成本。



 提示：


 1 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 数组 动态规划 👍 14 👎 0

*/

/*
题型: dp
题目: 最大化子数组的总成本
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumTotalCost(a []int) int64 {
	n := len(a)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return a[0]
		}
		p := &memo[i]
		if *p != math.MinInt {
			return *p
		}
		*p = max(dfs(i-1)+a[i], dfs(i-2)+a[i-1]-a[i])
		return *p
	}
	return int64(dfs(n - 1))
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	dfs(i)=max(dfs(i−1)+a[i],dfs(i−2)+a[i−1]−a[i])
*/

/*
递推
func maximumTotalCost(a []int) int64 {
	n := len(a)
	f := make([]int, n+1)
	f[1] = a[0]
	for i := 1; i < n; i++ {
		f[i+1] = max(f[i]+a[i], f[i-1]+a[i-1]-a[i])
	}
	return int64(f[n])
}

*/

/*
空间优化

	func maximumTotalCost(a []int) int64 {
		f0, f1 := 0, a[0]
		for i := 1; i < len(a); i++ {
			f0, f1 = f1, max(f1+a[i], f0+a[i-1]-a[i])
		}
		return int64(f1)
	}
*/
func main() {

}
