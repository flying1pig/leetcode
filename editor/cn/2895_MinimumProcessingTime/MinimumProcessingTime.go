package main

import "slices"

/**
你有 n 颗处理器，每颗处理器都有 4 个核心。现有 n * 4 个待执行任务，每个核心只执行 一次 任务。

 给你一个下标从 0 开始的整数数组 processorTime ，表示每颗处理器最早空闲时间。另给你一个下标从 0 开始的整数数组 tasks ，表示执行每个
任务所需的时间。返回所有任务都执行完毕需要的 最小时间 。

 注意：每个核心独立执行任务。



 示例 1：


输入：processorTime = [8,10], tasks = [2,2,3,1,8,7,4,5]
输出：16
解释：
最优的方案是将下标为 4, 5, 6, 7 的任务分配给第一颗处理器（最早空闲时间 time = 8），下标为 0, 1, 2, 3 的任务分配给第二颗处理器（
最早空闲时间 time = 10）。
第一颗处理器执行完所有任务需要花费的时间 = max(8 + 8, 8 + 7, 8 + 4, 8 + 5) = 16 。
第二颗处理器执行完所有任务需要花费的时间 = max(10 + 2, 10 + 2, 10 + 3, 10 + 1) = 13 。
因此，可以证明执行完所有任务需要花费的最小时间是 16 。

 示例 2：


输入：processorTime = [10,20], tasks = [2,3,1,2,5,8,4,3]
输出：23
解释：
最优的方案是将下标为 1, 4, 5, 6 的任务分配给第一颗处理器（最早空闲时间 time = 10），下标为 0, 2, 3, 7 的任务分配给第二颗处理器
（最早空闲时间 time = 20）。
第一颗处理器执行完所有任务需要花费的时间 = max(10 + 3, 10 + 5, 10 + 8, 10 + 4) = 18 。
第二颗处理器执行完所有任务需要花费的时间 = max(20 + 2, 20 + 1, 20 + 2, 20 + 3) = 23 。
因此，可以证明执行完所有任务需要花费的最小时间是 23 。




 提示：


 1 <= n == processorTime.length <= 25000
 1 <= tasks.length <= 10⁵
 0 <= processorTime[i] <= 10⁹
 1 <= tasks[i] <= 10⁹
 tasks.length == 4 * n


 Related Topics 贪心 数组 排序 👍 15 👎 0

*/

/*
题型: 贪心
题目: 最小处理时间
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minProcessingTime(processorTime, tasks []int) (ans int) {
	slices.Sort(processorTime)
	slices.SortFunc(tasks, func(a, b int) int { return b - a })
	for i, p := range processorTime {
		ans = max(ans, p+tasks[i*4])
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
