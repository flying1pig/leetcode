package main

import "container/heap"

/**
给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。



 示例 1:


输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]


 示例 2:


输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]




 提示:


 1 <= nums1.length, nums2.length <= 10⁵
 -10⁹ <= nums1[i], nums2[i] <= 10⁹
 nums1 和 nums2 均为 升序排列

 1 <= k <= 10⁴
 k <= nums1.length * nums2.length


 Related Topics 数组 堆（优先队列） 👍 655 👎 0

*/

/*
题型: 堆
题目: 查找和最小的 K 对数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	ans := make([][]int, 0, k) // 预分配空间
	h := make(hp, min(k, n))
	for i := range h {
		h[i] = tuple{nums1[i] + nums2[0], i, 0}
	}
	for len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
