package main

import (
	"container/heap"
	"sort"
)

/**
给你一个整数数组 nums ，一个整数 k 和一个整数 multiplier 。

 你需要对 nums 执行 k 次操作，每次操作中：


 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
 将 x 替换为 x * multiplier 。


 请你返回执行完 k 次乘运算之后，最终的 nums 数组。



 示例 1：


 输入：nums = [2,1,3,5,6], k = 5, multiplier = 2


 输出：[8,4,6,5,6]

 解释：




 操作
 结果


 1 次操作后
 [2, 2, 3, 5, 6]


 2 次操作后
 [4, 2, 3, 5, 6]


 3 次操作后
 [4, 4, 3, 5, 6]


 4 次操作后
 [4, 4, 6, 5, 6]


 5 次操作后
 [8, 4, 6, 5, 6]




 示例 2：


 输入：nums = [1,2], k = 3, multiplier = 4


 输出：[16,8]

 解释：




 操作
 结果


 1 次操作后
 [4, 2]


 2 次操作后
 [4, 8]


 3 次操作后
 [16, 8]






 提示：


 1 <= nums.length <= 100
 1 <= nums[i] <= 100
 1 <= k <= 10
 1 <= multiplier <= 5


 Related Topics 数组 数学 模拟 堆（优先队列） 👍 26 👎 0

*/

/*
题型: 堆
题目: K 次乘运算后的最终数组 I
*/

// leetcode submit region begin(Prohibit modification and deletion)
func quickMul(x, y, m int64) int64 {
	res := int64(1)
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % m
		}
		y >>= 1
		x = (x * x) % m
	}
	return res
}

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n, m := len(nums), int64(1e9+7)
	mx := 0
	var v minHeap
	for i, num := range nums {
		mx = max(mx, num)
		v = append(v, pair{int64(num), i})
	}
	heap.Init(&v)
	for ; v[0].first < int64(mx) && k > 0; k-- {
		x := heap.Pop(&v).(pair)
		x.first *= int64(multiplier)
		heap.Push(&v, x)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].first < v[j].first || v[i].first == v[j].first && v[i].second < v[j].second
	})
	for i := 0; i < n; i++ {
		t := k / n
		if i < k%n {
			t++
		}
		nums[v[i].second] = int((v[i].first % m) * quickMul(int64(multiplier), int64(t), m) % m)
	}
	return nums
}

type pair struct {
	first  int64
	second int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].first < h[j].first || h[i].first == h[j].first && h[i].second < h[j].second
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路: 对 nums 进行模拟操作，每次操作先找到 nums 的最前面的最小值，
然后将该元素替换成乘以 multiplier 后的值，最后返回 k 次模拟操作后的数组 nums

func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        m := 0
        for j := range nums {
            if nums[j] < nums[m] {
                m = j
            }
        }
        nums[m] *= multiplier
    }
    return nums
}

*/

func main() {

}
