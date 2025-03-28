package main

import (
	"container/heap"
	"sort"
)

/**
有 n 个朋友在举办一个派对，这些朋友从 0 到 n - 1 编号。派对里有 无数 张椅子，编号为 0 到 infinity 。当一个朋友到达派对时，他会占据
编号最小 且未被占据的椅子。


 比方说，当一个朋友到达时，如果椅子 0 ，1 和 5 被占据了，那么他会占据 2 号椅子。


 当一个朋友离开派对时，他的椅子会立刻变成未占据状态。如果同一时刻有另一个朋友到达，可以立即占据这张椅子。

 给你一个下标从 0 开始的二维整数数组 times ，其中 times[i] = [arrivali, leavingi] 表示第 i 个朋友到达和离开的时刻
，同时给你一个整数 targetFriend 。所有到达时间 互不相同 。

 请你返回编号为 targetFriend 的朋友占据的 椅子编号 。



 示例 1：

 输入：times = [[1,4],[2,3],[4,6]], targetFriend = 1
输出：1
解释：
- 朋友 0 时刻 1 到达，占据椅子 0 。
- 朋友 1 时刻 2 到达，占据椅子 1 。
- 朋友 1 时刻 3 离开，椅子 1 变成未占据。
- 朋友 0 时刻 4 离开，椅子 0 变成未占据。
- 朋友 2 时刻 4 到达，占据椅子 0 。
朋友 1 占据椅子 1 ，所以返回 1 。


 示例 2：

 输入：times = [[3,10],[1,5],[2,6]], targetFriend = 0
输出：2
解释：
- 朋友 1 时刻 1 到达，占据椅子 0 。
- 朋友 2 时刻 2 到达，占据椅子 1 。
- 朋友 0 时刻 3 到达，占据椅子 2 。
- 朋友 1 时刻 5 离开，椅子 0 变成未占据。
- 朋友 2 时刻 6 离开，椅子 1 变成未占据。
- 朋友 0 时刻 10 离开，椅子 2 变成未占据。
朋友 0 占据椅子 2 ，所以返回 2 。




 提示：


 n == times.length
 2 <= n <= 10⁴
 times[i].length == 2
 1 <= arrivali < leavingi <= 10⁵
 0 <= targetFriend <= n - 1
 每个 arrivali 时刻 互不相同 。


 Related Topics 数组 哈希表 堆（优先队列） 👍 53 👎 0

*/

/*
题型: 堆
题目: 最小未被占据椅子的编号
*/

// leetcode submit region begin(Prohibit modification and deletion)
func smallestChair(times [][]int, targetFriend int) int {
	// 按时间顺序，记录每个到达事件和离开事件相对应的朋友编号
	events := make([][2][]int, 1e5+1)
	for i, t := range times {
		l, r := t[0], t[1]
		events[l][1] = append(events[l][1], i) // 朋友到达
		events[r][0] = append(events[r][0], i) // 朋友离开
	}

	// 初始化未被占据的椅子
	n := len(times)
	unoccupied := hp{make([]int, n)}
	for i := range unoccupied.IntSlice {
		unoccupied.IntSlice[i] = i
	}

	// 按时间顺序扫描每个事件
	belong := make([]int, n)
	for _, e := range events {
		for _, id := range e[0] { // 朋友离开
			heap.Push(&unoccupied, belong[id]) // 返还椅子
		}
		for _, id := range e[1] { // 朋友到达
			belong[id] = heap.Pop(&unoccupied).(int) // 记录占据该椅子的朋友编号
			if id == targetFriend {
				return belong[id]
			}
		}
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
