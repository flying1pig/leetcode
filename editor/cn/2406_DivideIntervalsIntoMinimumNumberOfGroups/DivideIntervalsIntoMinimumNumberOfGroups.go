package main

import (
	"container/heap"
	"sort"
)

/**
给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti, righti]
。

 你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。

 请你返回 最少 需要划分成多少个组。

 如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。



 示例 1：


输入：intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
输出：3
解释：我们可以将区间划分为如下的区间组：
- 第 1 组：[1, 5] ，[6, 8] 。
- 第 2 组：[2, 3] ，[5, 10] 。
- 第 3 组：[1, 10] 。
可以证明无法将区间划分为少于 3 个组。


 示例 2：


输入：intervals = [[1,3],[5,6],[8,10],[11,13]]
输出：1
解释：所有区间互不相交，所以我们可以把它们全部放在一个组内。



 提示：


 1 <= intervals.length <= 10⁵
 intervals[i].length == 2
 1 <= lefti <= righti <= 10⁶


 Related Topics 贪心 数组 双指针 前缀和 排序 堆（优先队列） 👍 85 👎 0

*/

/*
题型: 堆
题目: 将区间分为最少组数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minGroups(intervals [][]int) int {
	h := hp{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, p := range intervals {
		if h.Len() == 0 || p[0] <= h.IntSlice[0] {
			heap.Push(&h, p[1])
		} else {
			h.IntSlice[0] = p[1]
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
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
