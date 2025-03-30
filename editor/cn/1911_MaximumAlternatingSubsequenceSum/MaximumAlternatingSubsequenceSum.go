package main

import "math"

/**
一个下标从 0 开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和 。


 比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4 。


 给你一个数组 nums ，请你返回 nums 中任意子序列的 最大交替和 （子序列的下标 重新 从 0 开始编号）。




 一个数组的 子序列 是从原数组中删除一些元素后（也可能一个也不删除）剩余元素不改变顺序组成的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4]
的一个子序列（加粗元素），但是 [2,4,2] 不是。



 示例 1：

 输入：nums = [4,2,5,3]
输出：7
解释：最优子序列为 [4,2,5] ，交替和为 (4 + 5) - 2 = 7 。


 示例 2：

 输入：nums = [5,6,7,8]
输出：8
解释：最优子序列为 [8] ，交替和为 8 。


 示例 3：

 输入：nums = [6,2,1,2,4,5]
输出：10
解释：最优子序列为 [6,1,5] ，交替和为 (6 + 5) - 1 = 10 。




 提示：


 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁵


 Related Topics 数组 动态规划 👍 195 👎 0

*/

/*
题型: dp
题目: 最大交替子序列和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxAlternatingSum(nums []int) int64 {
	f := [2]int{0, math.MinInt64 / 2} // 除 2 防止计算时溢出
	for _, v := range nums {
		f = [2]int{max(f[0], f[1]-v), max(f[1], f[0]+v)}
	}
	return int64(f[1])
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	f[i+1][0]=max(f[i][0],f[i][1]−nums[i])
	f[i+1][1]=max(f[i][1],f[i][0]+nums[i])

*/

func main() {

}
