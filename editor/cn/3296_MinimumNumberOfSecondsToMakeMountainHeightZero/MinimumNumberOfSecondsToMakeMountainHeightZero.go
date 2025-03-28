package main

import "container/heap"

/**
给你一个整数 mountainHeight 表示山的高度。

 同时给你一个整数数组 workerTimes，表示工人们的工作时间（单位：秒）。

 工人们需要 同时 进行工作以 降低 山的高度。对于工人 i :


 山的高度降低 x，需要花费 workerTimes[i] + workerTimes[i] * 2 + ... + workerTimes[i] * x 秒。
例如：



 山的高度降低 1，需要 workerTimes[i] 秒。
 山的高度降低 2，需要 workerTimes[i] + workerTimes[i] * 2 秒，依此类推。




 返回一个整数，表示工人们使山的高度降低到 0 所需的 最少 秒数。



 示例 1：


 输入： mountainHeight = 4, workerTimes = [2,1,1]


 输出： 3

 解释：

 将山的高度降低到 0 的一种方式是：


 工人 0 将高度降低 1，花费 workerTimes[0] = 2 秒。
 工人 1 将高度降低 2，花费 workerTimes[1] + workerTimes[1] * 2 = 3 秒。
 工人 2 将高度降低 1，花费 workerTimes[2] = 1 秒。


 因为工人同时工作，所需的最少时间为 max(2, 3, 1) = 3 秒。

 示例 2：


 输入： mountainHeight = 10, workerTimes = [3,2,2,4]


 输出： 12

 解释：


 工人 0 将高度降低 2，花费 workerTimes[0] + workerTimes[0] * 2 = 9 秒。
 工人 1 将高度降低 3，花费 workerTimes[1] + workerTimes[1] * 2 + workerTimes[1] * 3 = 12 秒
。
 工人 2 将高度降低 3，花费 workerTimes[2] + workerTimes[2] * 2 + workerTimes[2] * 3 = 12 秒
。
 工人 3 将高度降低 2，花费 workerTimes[3] + workerTimes[3] * 2 = 12 秒。


 所需的最少时间为 max(9, 12, 12, 12) = 12 秒。

 示例 3：


 输入： mountainHeight = 5, workerTimes = [1]


 输出： 15

 解释：

 这个示例中只有一个工人，所以答案是 workerTimes[0] + workerTimes[0] * 2 + workerTimes[0] * 3 +
workerTimes[0] * 4 + workerTimes[0] * 5 = 15 秒。



 提示：


 1 <= mountainHeight <= 10⁵
 1 <= workerTimes.length <= 10⁴
 1 <= workerTimes[i] <= 10⁶


 Related Topics 贪心 数组 数学 二分查找 堆（优先队列） 👍 20 👎 0

*/

/*
题型: 堆
题目: 移山所需的最少秒数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	h := make(hp, len(workerTimes))
	for i, t := range workerTimes {
		h[i] = worker{t, t, t}
	}
	heap.Init(&h)

	ans := 0
	for ; mountainHeight > 0; mountainHeight-- {
		ans = h[0].nxt // 最后一个出堆的 nxt 即为答案
		h[0].delta += h[0].base
		h[0].nxt += h[0].delta
		heap.Fix(&h, 0)
	}
	return int64(ans)
}

// 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
type worker struct{ nxt, delta, base int }
type hp []worker

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].nxt < h[j].nxt }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
