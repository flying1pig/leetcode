package main

import "fmt"

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



 示例 1：


输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

 示例 2：


输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。




 提示：


 1 <= nums.length <= 200
 1 <= nums[i] <= 100


 Related Topics 数组 动态规划 👍 2274 👎 0

*/

/*
题型: dp
题目: 分割等和子集
*/

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	s := 0
	for _, val := range nums {
		s += val
	}
	if s%2 == 1 {
		return false
	}
	target := s / 2
	n := len(nums)

	mem := make([][]int8, n)
	for i := range mem {
		mem[i] = make([]int8, target+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var f func(int, int) bool
	f = func(i int, j int) bool {
		if i < 0 {
			return j == 0
		}
		if mem[i][j] != -1 {
			return mem[i][j] == 1
		}

		if j < nums[i] {
			if f(i-1, j) {
				mem[i][j] = 1
			} else {
				mem[i][j] = 0
			}
			return mem[i][j] == 1
		}

		if f(i-1, j) || f(i-1, j-nums[i]) {
			mem[i][j] = 1
		} else {
			mem[i][j] = 0
		}
		return mem[i][j] == 1
	}
	return f(n-1, target)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	arr := make([]bool, 5)
	fmt.Println(arr)
}
