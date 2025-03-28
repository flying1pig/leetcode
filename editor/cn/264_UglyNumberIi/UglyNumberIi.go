package main

import (
	"container/heap"
	"sort"
)

/**
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

 丑数 就是质因子只包含 2、3 和 5 的正整数。



 示例 1：


输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。


 示例 2：


输入：n = 1
输出：1
解释：1 通常被视为丑数。




 提示：


 1 <= n <= 1690


 Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 1252 👎 0

*/

/*
题型: 堆
题目: 丑数 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
var factors = []int{2, 3, 5}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
