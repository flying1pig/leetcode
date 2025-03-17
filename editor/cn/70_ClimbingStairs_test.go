package leet

import (
	"fmt"
	"testing"
)

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



 示例 1：


输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

 示例 2：


输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶




 提示：


 1 <= n <= 45


 👍 3764 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	f, f_1, f_2 := 1, 1, 1

	for i := 2; i <= n; i++ {
		newf := f_1 + f_2
		f_2 = f_1
		f_1 = newf
		f = newf
	}
	return f
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: f(i) = f(i-1) + f(i-2)
边界条件: f(1) = 0, f(0) = 0
*/

/*
递归:
func climbStairs(n int) int {
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		return dfs(i-1) + dfs(i-2)
	}
	return dfs(n)
}
时间复杂度：o(n^2)
空间复杂度：o(n)
*/

/*
记忆化搜索:
func climbStairs(n int) int {
	var dfs func(int) int
	temp := make([]int,n+1)
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		if temp[i] == 0 {
			temp[i] = dfs(i-1)+dfs(i-2)
		}
		return temp[i]
	}
	return dfs(n)
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func climbStairs(n int) int {
	mem := make([]int,n+1)
	mem[0],mem[1]=1,1
	for i := 2;i<=n;i++ {
		mem[i] = mem[i-1]+mem[i-2]
	}
	return mem[n]
}
时间复杂度: o(n)
空间复杂度: o(n)

空间复杂度优化:
func climbStairs(n int) int {
	f, f_1, f_2 := 1, 1, 1

	for i := 2;i<=n;i++ {
		newf := f_1+f_2
		f_2 = f_1
		f_1 = newf
		f = newf
	}
	return f
}
时间复杂度: o(n)
空间复杂度: o(1)
*/

func TestClimbingStairs(t *testing.T) {
	fmt.Println(climbStairs(3))

}
