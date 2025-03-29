package main

/**
给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。



 示例 1：


输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]


 示例 2：


输入：nums = [4,4,3,2,1]
输出：[[4,4]]




 提示：


 1 <= nums.length <= 15
 -100 <= nums[i] <= 100


 Related Topics 位运算 数组 哈希表 回溯 👍 854 👎 0

*/

/*
题型: 回溯
题目: 非递减子序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(index int, lis []int) {
		if len(lis) >= 2 {
			dest := make([]int, len(lis))
			copy(dest, lis)
			res = append(res, dest)
		}
		//保证一层循环中不会选择重复的值就ok了，画一下可能会比较清晰
		var visit [201]bool
		for i := index; i < len(nums); i++ {
			if visit[nums[i]+100] {
				continue
			}
			if len(lis) == 0 || nums[i] >= lis[len(lis)-1] {
				visit[nums[i]+100] = true
				dfs(i+1, append(lis, nums[i]))
			}
		}
	}
	dfs(0, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
