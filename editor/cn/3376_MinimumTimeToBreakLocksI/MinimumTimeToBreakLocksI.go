package main

import "math"

/**
Bob 被困在了一个地窖里，他需要破解 n 个锁才能逃出地窖，每一个锁都需要一定的 能量 才能打开。每一个锁需要的能量存放在一个数组 strength 里，其中
 strength[i] 表示打开第 i 个锁需要的能量。

 Bob 有一把剑，它具备以下的特征：


 一开始剑的能量为 0 。
 剑的能量增加因子 X 一开始的值为 1 。
 每分钟，剑的能量都会增加当前的 X 值。
 打开第 i 把锁，剑的能量需要到达 至少 strength[i] 。
 打开一把锁以后，剑的能量会变回 0 ，X 的值会增加一个给定的值 K 。


 你的任务是打开所有 n 把锁并逃出地窖，请你求出需要的 最少 分钟数。

 请你返回 Bob 打开所有 n 把锁需要的 最少 时间。



 示例 1：


 输入：strength = [3,4,1], K = 1


 输出：4

 解释：




 时间
 能量
 X
 操作
 更新后的 X


 0
 0
 1
 什么也不做
 1


 1
 1
 1
 打开第 3 把锁
 2


 2
 2
 2
 什么也不做
 2


 3
 4
 2
 打开第 2 把锁
 3


 4
 3
 3
 打开第 1 把锁
 3




 无法用少于 4 分钟打开所有的锁，所以答案为 4 。

 示例 2：


 输入：strength = [2,5,4], K = 2


 输出：5

 解释：




 时间
 能量
 X
 操作
 更新后的 X


 0
 0
 1
 什么也不做
 1


 1
 1
 1
 什么也不做
 1


 2
 2
 1
 打开第 1 把锁
 3


 3
 3
 3
 什么也不做
 3


 4
 6
 3
 打开第 2 把锁
 5


 5
 5
 5
 打开第 3 把锁
 7




 无法用少于 5 分钟打开所有的锁，所以答案为 5 。



 提示：


 n == strength.length
 1 <= n <= 8
 1 <= K <= 10
 1 <= strength[i] <= 10⁶


 Related Topics 位运算 深度优先搜索 数组 动态规划 回溯 状态压缩 👍 8 👎 0

*/

/*
题型: 回溯
题目: 破解锁的最少时间 I
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findMinimumTime(strength []int, k int) int {
	ans := math.MaxInt
	n := len(strength)
	done := make([]bool, n)
	var dfs func(int, int)
	dfs = func(i, time int) {
		// 最优性剪枝：答案不可能变小
		if time >= ans {
			return
		}
		if i == n {
			ans = time
			return
		}
		x := 1 + k*i
		for j, d := range done {
			if !d {
				done[j] = true // 已开锁
				dfs(i+1, time+(strength[j]-1)/x+1)
				done[j] = false // 恢复现场
			}
		}
	}
	dfs(0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
