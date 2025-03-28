package main

import (
	"container/heap"
	"slices"
	"sort"
)

/**
给你两个整数数组，nums1 和 nums2，长度均为 n，以及一个正整数 k 。

 对从 0 到 n - 1 每个下标 i ，执行下述操作：


 找出所有满足 nums1[j] 小于 nums1[i] 的下标 j 。
 从这些下标对应的 nums2[j] 中选出 至多 k 个，并 最大化 这些值的总和作为结果。


 返回一个长度为 n 的数组 answer ，其中 answer[i] 表示对应下标 i 的结果。



 示例 1：


 输入：nums1 = [4,2,1,5,3], nums2 = [10,20,30,40,50], k = 2


 输出：[80,30,0,80,50]

 解释：


 对于 i = 0 ：满足 nums1[j] < nums1[0] 的下标为 [1, 2, 4] ，选出其中值最大的两个，结果为 50 + 30 = 80 。

 对于 i = 1 ：满足 nums1[j] < nums1[1] 的下标为 [2] ，只能选择这个值，结果为 30 。
 对于 i = 2 ：不存在满足 nums1[j] < nums1[2] 的下标，结果为 0 。
 对于 i = 3 ：满足 nums1[j] < nums1[3] 的下标为 [0, 1, 2, 4] ，选出其中值最大的两个，结果为 50 + 30 = 80
 。
 对于 i = 4 ：满足 nums1[j] < nums1[4] 的下标为 [1, 2] ，选出其中值最大的两个，结果为 30 + 20 = 50 。


 示例 2：


 输入：nums1 = [2,2,2,2], nums2 = [3,1,2,3], k = 1


 输出：[0,0,0,0]

 解释：由于 nums1 中的所有元素相等，不存在满足条件 nums1[j] < nums1[i]，所有位置的结果都是 0 。



 提示：


 n == nums1.length == nums2.length
 1 <= n <= 10⁵
 1 <= nums1[i], nums2[i] <= 10⁶
 1 <= k <= n


 Related Topics 数组 排序 堆（优先队列） 👍 8 👎 0

*/

/*
题型: 堆
题目: 选出和最大的 K 个元素
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	// 分组循环模板
	for i := 0; i < n; {
		start := i
		// 找到所有相同的 nums1[i]，这些数的答案都是一样的
		x := a[start].x
		for ; i < n && a[i].x == x; i++ {
			ans[a[i].i] = int64(s)
		}
		// 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
		for ; start < i; start++ {
			y := a[start].y
			s += y
			heap.Push(h, y)
			if h.Len() > k {
				s -= heap.Pop(h).(int)
			}
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
