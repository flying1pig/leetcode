package main

import "math"

/**
给你一个大小为 4 的整数数组 a 和一个大小 至少为 4 的整数数组 b。

 你需要从数组 b 中选择四个下标 i0, i1, i2, 和 i3，并满足 i0 < i1 < i2 < i3。你的得分将是 a[0] * b[i0] +
a[1] * b[i1] + a[2] * b[i2] + a[3] * b[i3] 的值。

 返回你能够获得的 最大 得分。



 示例 1：


 输入： a = [3,2,5,6], b = [2,-6,4,-5,-3,2,-7]


 输出： 26

 解释： 选择下标 0, 1, 2 和 5。得分为 3 * 2 + 2 * (-6) + 5 * 4 + 6 * 2 = 26。

 示例 2：


 输入： a = [-1,4,5,-2], b = [-5,-1,-3,-2,-4]


 输出： -1

 解释： 选择下标 0, 1, 3 和 4。得分为 (-1) * (-5) + 4 * (-1) + 5 * (-2) + (-2) * (-4) = -1。




 提示：


 a.length == 4
 4 <= b.length <= 10⁵
 -10⁵ <= a[i], b[i] <= 10⁵


 Related Topics 数组 动态规划 👍 16 👎 0

*/

/*
题型: dp
题目: 最高乘法得分
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(a, b []int) int64 {
	n := len(b)
	memo := make([][4]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = math.MinInt64 // 表示没有计算过
		}
	}
	var dfs func(int, int) int64
	dfs = func(i, j int) int64 {
		if j < 0 { // 选完了
			return 0
		}
		if i < 0 { // j >= 0，没选完
			return math.MinInt64 / 2 // 防止溢出
		}
		p := &memo[i][j]
		if *p == math.MinInt64 { // 需要计算，并记忆化
			*p = max(dfs(i-1, j), dfs(i-1, j-1)+int64(a[j])*int64(b[i]))
		}
		return *p
	}
	return dfs(n-1, 3)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
	dfs(i,j)=max(dfs(i−1,j),dfs(i−1,j−1)+a[j]⋅b[i])
边界条件:
	dfs(i,−1)=0, dfs(−1,j≥0)= -inf
*/

func main() {

}
