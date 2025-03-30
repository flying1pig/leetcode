package main

import (
	"fmt"
	"math"
)

/**
给你一个下标从 0 开始的整数数组 nums 和一个整数 target 。

 返回和为 target 的 nums 子序列中，子序列 长度的最大值 。如果不存在和为 target 的子序列，返回 -1 。

 子序列 指的是从原数组中删除一些或者不删除任何元素后，剩余元素保持原来的顺序构成的数组。



 示例 1：


输入：nums = [1,2,3,4,5], target = 9
输出：3
解释：总共有 3 个子序列的和为 9 ：[4,5] ，[1,3,5] 和 [2,3,4] 。最长的子序列是 [1,3,5] 和 [2,3,4] 。所以答案为 3
 。


 示例 2：


输入：nums = [4,1,3,2,1,5], target = 7
输出：4
解释：总共有 5 个子序列的和为 7 ：[4,3] ，[4,1,2] ，[4,2,1] ，[1,1,5] 和 [1,3,2,1] 。最长子序列为 [1,3,2,
1] 。所以答案为 4 。


 示例 3：


输入：nums = [1,1,5,4,5], target = 3
输出：-1
解释：无法得到和为 3 的子序列。




 提示：


 1 <= nums.length <= 1000
 1 <= nums[i] <= 1000
 1 <= target <= 1000


 Related Topics 数组 动态规划 👍 56 👎 0

*/

/*
题型: dp
题目: 和为目标值的最长子序列的长度
*/

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	ans := -1

	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, target+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var f func(int, int) int
	f = func(i int, j int) int {
		if i < 0 {
			if j == 0 {
				return 0
			}
			return math.MinInt
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		if j < nums[i] {
			mem[i][j] = f(i-1, j)
			return mem[i][j]
		}

		mem[i][j] = max(f(i-1, j), f(i-1, j-nums[i])+1)
		ans = max(ans, mem[i][j])
		return mem[i][j]
	}

	f(n-1, target)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)来计算所有和为目标值方案的最大长度
f(i,j) = max(f(i-1,j), f(i-1,j-nums[i]))
*/

/*
记忆化搜索:
*/
func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(lengthOfLongestSubsequence(arr, 9))
}
