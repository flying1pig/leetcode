package main

import "slices"

/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



 示例 1：


输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]


 示例 2：


输入：nums = [0]
输出：[[],[0]]




 提示：


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10
 nums 中的所有元素 互不相同


 Related Topics 位运算 数组 回溯 👍 2460 👎 0

*/

/*
题型: 回溯
题目: 子集
*/

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n) // 预分配空间
	path := make([]int, 0, n)     // 预分配空间
	var dfs func(int)
	dfs = func(i int) {
		if i == n { // 子集构造完毕
			ans = append(ans, slices.Clone(path)) // 复制 path
			return
		}

		// 不选 nums[i]
		dfs(i + 1)

		// 选 nums[i]
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
从答案的视角看，枚举选哪个
func subsets(nums []int) [][]int {
    n := len(nums)
    ans := make([][]int, 0, 1<<n) // 预分配空间
    path := make([]int, 0, n) // 预分配空间
    var dfs func(int)
    dfs = func(i int) {
        ans = append(ans, slices.Clone(path)) // 复制 path
        for j := i; j < n; j++ { // 枚举选择的数字
            path = append(path, nums[j])
            dfs(j + 1)
            path = path[:len(path)-1] // 恢复现场
        }
    }
    dfs(0)
    return ans
}

*/

func main() {

}
