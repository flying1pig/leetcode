package main

import (
	"container/heap"
	"sort"
)

/**
给你一个下标从 0 开始的二维整数数组 events ，其中 events[i] = [startTimei, endTimei, valuei] 。第 i 个
活动开始于 startTimei ，结束于 endTimei ，如果你参加这个活动，那么你可以得到价值 valuei 。你 最多 可以参加 两个时间不重叠 活动
，使得它们的价值之和 最大 。

 请你返回价值之和的 最大值 。

 注意，活动的开始时间和结束时间是 包括 在活动时间内的，也就是说，你不能参加两个活动且它们之一的开始时间等于另一个活动的结束时间。更具体的，如果你参加一个活动
，且结束时间为 t ，那么下一个活动必须在 t + 1 或之后的时间开始。



 示例 1:



 输入：events = [[1,3,2],[4,5,2],[2,4,3]]
输出：4
解释：选择绿色的活动 0 和 1 ，价值之和为 2 + 2 = 4 。


 示例 2：



 输入：events = [[1,3,2],[4,5,2],[1,5,5]]
输出：5
解释：选择活动 2 ，价值和为 5 。


 示例 3：



 输入：events = [[1,5,3],[1,5,1],[6,6,5]]
输出：8
解释：选择活动 0 和 2 ，价值之和为 3 + 5 = 8 。



 提示：


 2 <= events.length <= 10⁵
 events[i].length == 3
 1 <= startTimei <= endTimei <= 10⁹
 1 <= valuei <= 10⁶


 Related Topics 数组 二分查找 动态规划 排序 堆（优先队列） 👍 62 👎 0

*/

/*
题型: dp
题目: 两个最好的不重叠活动
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxTwoEvents(events [][]int) (ans int) {
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] }) // 按开始时间排序
	h := hp{}
	maxVal := 0
	for _, e := range events {
		start, end, val := e[0], e[1], e[2]
		for len(h) > 0 && h[0].end < start { // 如果结束时间早于当前活动开始时间
			maxVal = max(maxVal, heap.Pop(&h).(pair).val) // 更新前面可以选择的活动的最大价值
		}
		ans = max(ans, maxVal+val) // 至多参加两个活动
		heap.Push(&h, pair{end, val})
	}
	return
}

// heap 模板
type pair struct{ end, val int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
