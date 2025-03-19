package leet

import (
	"testing"
)

/**
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

 请你计算并返回达到楼梯顶部的最低花费。



 示例 1：


输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。


 示例 2：


输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。




 提示：


 2 <= cost.length <= 1000
 0 <= cost[i] <= 999


 👍 1629 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f1, f2 := 0, 0
	for i := 2; i <= n; i++ {
		newf := min(f1+cost[i-1], f2+cost[i-2])
		f2 = f1
		f1 = newf
	}
	return f1
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: f(i) = min(f(i-1) + cost[i-1], f(i-2) + cost[i-2])
边界条件: f(0) = 0 f(1) = 0
*/

/*
递归:
func minCostClimbingStairs(cost []int) int {
    var dfs func(i int) int
    dfs = func(i int) int {
        if i <= 1 {
            return 0
        }
        return min(dfs(i-1) + cost[i-1],dfs(i-2) + cost[i-2])
    }
    return dfs(len(cost))
}
时间复杂度: o(n^2)
空间复杂度: o(n)
*/

/*
记忆化搜索:
func minCostClimbingStairs(cost []int) int {
    var dfs func(i int) int
    n := len(cost)
    mem := make([]int,n+1)
    for i := range mem {
        mem[i] = -1
    }
    dfs = func(i int) int {
        if i <= 1 {
            return 0
        }
        if mem[i] == -1 {
            mem[i] = min(dfs(i-1) + cost[i-1],dfs(i-2) + cost[i-2])
        }
        return mem[i]
    }
    return dfs(len(cost))
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    mem := make([]int,n+1)
    for i:=2;i<=n;i++ {
        mem[i] = min(mem[i-1]+cost[i-1],mem[i-2]+cost[i-2])
    }
    return mem[n]
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推空间优化:
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    f1,f2 := 0,0
    for i:=2;i<=n;i++ {
        newf := min(f1+cost[i-1],f2+cost[i-2])
        f2 = f1
        f1 = newf
    }
    return f1
}
时间复杂度: o(n)
空间复杂度: o(1)
*/

func TestMinCostClimbingStairs(t *testing.T) {

}
