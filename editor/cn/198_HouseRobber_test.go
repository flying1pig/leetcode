package leet

import (
	"fmt"
	"testing"
)

/**
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小
偷闯入，系统会自动报警。

 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



 示例 1：


输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

 示例 2：


输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。




 提示：


 1 <= nums.length <= 100
 0 <= nums[i] <= 400


 👍 3203 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	m0 := 0
	m1 := nums[0]
	for i := 2; i <= len(nums); i++ {
		newm := max(m1, m0+nums[i-1])
		m0 = m1
		m1 = newm
	}
	return m1
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: f(i) = max(f(i-1),f(i-2)+nums[i-1])
边界条件: f(i) = 0, i <= 0
*/

/*
递归:
func rob(nums []int) int {
    var f func(i int) int
    f = func(i int) int {
        if i <=0 {
            return 0
        }
        return max(f(i-1),f(i-2)+nums[i-1])
    }
    return f(len(nums))
}
时间复杂度: o(n^2)
空间复杂度: o(1)
*/

/*
记忆化搜索:
func rob(nums []int) int {
    var f func(i int) int
    mem := make([]int,len(nums)+1)
    for i := range mem {
        mem[i] = -1
    }
    f = func(i int) int {
        if i <=0 {
            return 0
        }
        if mem[i] == -1 {
            mem[i] = max(f(i-1),f(i-2)+nums[i-1])
        }
        return mem[i]
    }
    return f(len(nums))
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func rob(nums []int) int {
    mem := make([]int,len(nums)+1)
    mem[0] = 0
    mem[1] = nums[0]
    for i:=2;i<=len(nums);i++ {
        mem[i] = max(mem[i-1],mem[i-2]+nums[i-1])
    }
    return mem[len(nums)]
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
空间优化:
func rob(nums []int) int {
    m0 := 0
    m1 := nums[0]
    for i:=2;i<=len(nums);i++ {
        newm := max(m1,m0+nums[i-1])
        m0 = m1
        m1 = newm
    }
    return m1
}
时间复杂度: o(n)
空间复杂度: o(1)
*/

func TestHouseRobber(t *testing.T) {
	nums := []int{2, 1}
	fmt.Println(rob(nums))
}
