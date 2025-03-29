package main

import "slices"

/**
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



 示例 1：


输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]


 示例 2：


输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]




 提示：


 1 <= nums.length <= 8
 -10 <= nums[i] <= 10


 Related Topics 数组 回溯 排序 👍 1695 👎 0

*/

/*
题型: 回溯
题目: 全排列 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	path := make([]int, n)    // 所有排列的长度都是 n
	onPath := make([]bool, n) // onPath[j] 表示 nums[j] 是否已经填入排列
	var dfs func(int)
	dfs = func(i int) { // i 表示当前要填排列的第几个数
		if i == n { // 填完了
			ans = append(ans, slices.Clone(path))
			return
		}
		// 枚举 nums[j] 填入 path[i]
		for j, on := range onPath {
			// 如果 nums[j] 已填入排列，continue
			// 如果 nums[j] 和前一个数 nums[j-1] 相等，且 nums[j-1] 没填入排列，continue
			if on || j > 0 && nums[j] == nums[j-1] && !onPath[j-1] {
				continue
			}
			path[i] = nums[j] // 填入排列
			onPath[j] = true  // nums[j] 已填入排列（注意标记的是下标，不是值）
			dfs(i + 1)        // 填排列的下一个数
			onPath[j] = false // 恢复现场
			// 注意 path 无需恢复现场，因为排列长度固定，直接覆盖就行
		}
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
