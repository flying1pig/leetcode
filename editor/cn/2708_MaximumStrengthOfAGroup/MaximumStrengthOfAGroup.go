package main

/**
给你一个下标从 0 开始的整数数组 nums ，它表示一个班级中所有学生在一次考试中的成绩。老师想选出一部分同学组成一个 非空 小组，且这个小组的 实力值 最大
，如果这个小组里的学生下标为 i0, i1, i2, ... , ik ，那么这个小组的实力值定义为 nums[i0] * nums[i1] * nums[i2
] * ... * nums[ik] 。

 请你返回老师创建的小组能得到的最大实力值为多少。



 示例 1：

 输入：nums = [3,-1,-5,2,5,-9]
输出：1350
解释：一种构成最大实力值小组的方案是选择下标为 [0,2,3,4,5] 的学生。实力值为 3 * (-5) * 2 * 5 * (-9) = 1350 ，这是可
以得到的最大实力值。


 示例 2：

 输入：nums = [-4,-5,-4]
输出：20
解释：选择下标为 [0, 1] 的学生。得到的实力值为 20 。我们没法得到更大的实力值。




 提示：


 1 <= nums.length <= 13
 -9 <= nums[i] <= 9


 Related Topics 贪心 位运算 数组 动态规划 回溯 枚举 排序 👍 62 👎 0

*/

/*
题型: dp
题目: 一个小组的最大实力值
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxStrength(nums []int) int64 {
	mn, mx := nums[0], nums[0]
	for _, x := range nums[1:] {
		mn, mx = min(mn, x, mn*x, mx*x),
			max(mx, x, mn*x, mx*x)
	}
	return int64(mx)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
mn=min(mn,x,mn⋅x,mx⋅x)
mx=max(mx,x,mn⋅x,mx⋅x)
*/

func main() {

}
