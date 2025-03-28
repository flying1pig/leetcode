package main

import (
	"container/heap"
	"fmt"
)

/**
给你一个非负整数数组 nums 和一个整数 k 。每次操作，你可以选择 nums 中 任一 元素并将它 增加 1 。

 请你返回 至多 k 次操作后，能得到的 nums的 最大乘积 。由于答案可能很大，请你将答案对 10⁹ + 7 取余后返回。



 示例 1：

 输入：nums = [0,4], k = 5
输出：20
解释：将第一个数增加 5 次。
得到 nums = [5, 4] ，乘积为 5 * 4 = 20 。
可以证明 20 是能得到的最大乘积，所以我们返回 20 。
存在其他增加 nums 的方法，也能得到最大乘积。


 示例 2：

 输入：nums = [6,3,3,2], k = 2
输出：216
解释：将第二个数增加 1 次，将第四个数增加 1 次。
得到 nums = [6, 4, 3, 3] ，乘积为 6 * 4 * 3 * 3 = 216 。
可以证明 216 是能得到的最大乘积，所以我们返回 216 。
存在其他增加 nums 的方法，也能得到最大乘积。




 提示：


 1 <= nums.length, k <= 10⁵
 0 <= nums[i] <= 10⁶


 Related Topics 贪心 数组 堆（优先队列） 👍 38 👎 0

*/

/*
题型: 堆
题目: K 次增加后的最大乘积
*/

// leetcode submit region begin(Prohibit modification and deletion)
type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]

}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	h := hp{}
	h = nums
	fmt.Println(h)
	if len(nums) == 0 {
		return 0
	}
	heap.Init(&h)

	for k > 0 {

		k--
		h[0]++
		heap.Fix(&h, 0)
	}

	res := h[0]
	for _, j := range nums[1:] {
		res = res * j % 1000000007
	}

	fmt.Println(res)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
