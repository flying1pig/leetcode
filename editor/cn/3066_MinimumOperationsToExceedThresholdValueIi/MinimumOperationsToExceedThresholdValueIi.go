package main

import "container/heap"

/**
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

 你可以对 nums 执行一些操作，在一次操作中，你可以：


 选择 nums 中 最小 的两个整数 x 和 y 。
 将 x 和 y 从 nums 中删除。
 将 min(x, y) * 2 + max(x, y) 添加到数组中的任意位置。


 注意，只有当 nums 至少 包含两个元素时，你才可以执行以上操作。

 你需要使数组中的所有元素都 大于或等于 k ，请你返回需要的 最少 操作次数。



 示例 1：


 输入：nums = [2,11,10,1,3], k = 10


 输出：2

 解释：


 第一次操作中，我们删除元素 1 和 2 ，然后添加 1 * 2 + 2 到 nums 中，nums 变为 [4, 11, 10, 3] 。
 第二次操作中，我们删除元素 3 和 4 ，然后添加 3 * 2 + 4 到 nums 中，nums 变为 [10, 11, 10] 。


 此时，数组中的所有元素都大于等于 10 ，所以我们停止操作。

 可以证明使数组中所有元素都大于等于 10 需要的最少操作次数为 2 。



 示例 2：


 输入：nums = [1,1,2,4,9], k = 20


 输出：4

 解释：


 第一次操作后，nums 变为 [2, 4, 9, 3]。
 第二次操作后，nums 变为 [7, 4, 9]。
 第三次操作后，nums 变为 [15, 9]。
 第四次操作后，nums 变为 [33]。


 此时，nums 中的所有元素都大于等于 20 ，所以我们停止操作。

 可以证明使数组中所有元素都大于等于 20 需要的最少操作次数为 4 。



 提示：


 2 <= nums.length <= 2 * 10⁵
 1 <= nums[i] <= 10⁹
 1 <= k <= 10⁹
 输入保证答案一定存在，也就是说，在进行某些次数的操作后，数组中所有元素都大于等于 k 。


 Related Topics 数组 模拟 堆（优先队列） 👍 34 👎 0

*/

/*
题型: 堆
题目: 超过阈值的最少操作数 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, k int) int {
	res := 0
	pq := &MinHeap{}
	heap.Init(pq)
	for _, num := range nums {
		heap.Push(pq, num)
	}

	for (*pq)[0] < k {
		x := heap.Pop(pq).(int)
		y := heap.Pop(pq).(int)
		heap.Push(pq, x+x+y)
		res++
	}

	return res
}

// MinHeap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

//小顶堆

func main() {

}
