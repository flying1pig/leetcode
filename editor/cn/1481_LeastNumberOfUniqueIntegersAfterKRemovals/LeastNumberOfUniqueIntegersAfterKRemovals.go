package main

import "container/heap"

/**
给你一个整数数组 arr 和一个整数 k 。现需要从数组中恰好移除 k 个元素，请找出移除后数组中不同整数的最少数目。






 示例 1：

 输入：arr = [5,5,4], k = 1
输出：1
解释：移除 1 个 4 ，数组中只剩下 5 一种整数。


 示例 2：

 输入：arr = [4,3,1,1,3,3,2], k = 3
输出：2
解释：先移除 4、2 ，然后再移除两个 1 中的任意 1 个或者三个 3 中的任意 1 个，最后剩下 1 和 3 两种整数。



 提示：


 1 <= arr.length <= 10^5
 1 <= arr[i] <= 10^9
 0 <= k <= arr.length


 Related Topics 贪心 数组 哈希表 计数 排序 👍 128 👎 0

*/

/*
题型: 贪心
题目: 不同整数的最少数目
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findLeastNumOfUniqueInts(arr []int, k int) int {
	cntMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		cntMap[arr[i]]++
	}
	h := pairHeap{}
	for n, c := range cntMap {
		heap.Push(&h, pair{n, c})
	}
	for len(h) > 0 && h[0].cnt <= k {
		top := heap.Pop(&h).(pair)
		k -= top.cnt
	}
	return len(h)
}

type pair struct {
	num int
	cnt int
}

type pairHeap []pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(i, j int) bool  { return h[i].cnt < h[j].cnt }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
