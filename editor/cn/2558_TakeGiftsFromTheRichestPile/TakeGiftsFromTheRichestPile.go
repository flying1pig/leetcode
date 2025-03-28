package main

import (
	"container/heap"
	"math"
)

/**
给你一个整数数组 gifts ，表示各堆礼物的数量。每一秒，你需要执行以下操作：


 选择礼物数量最多的那一堆。
 如果不止一堆都符合礼物数量最多，从中选择任一堆即可。
 将堆中的礼物数量减少到堆中原来礼物数量的平方根，向下取整。


 返回在 k 秒后剩下的礼物数量。



 示例 1：


输入：gifts = [25,64,9,4,100], k = 4
输出：29
解释：
按下述方式取走礼物：
- 在第一秒，选中最后一堆，剩下 10 个礼物。
- 接着第二秒选中第二堆礼物，剩下 8 个礼物。
- 然后选中第一堆礼物，剩下 5 个礼物。
- 最后，再次选中最后一堆礼物，剩下 3 个礼物。
最后剩下的礼物数量分别是 [5,8,9,4,3] ，所以，剩下礼物的总数量是 29 。


 示例 2：


输入：gifts = [1,1,1,1], k = 4
输出：4
解释：
在本例中，不管选中哪一堆礼物，都必须剩下 1 个礼物。
也就是说，你无法获取任一堆中的礼物。
所以，剩下礼物的总数量是 4 。




 提示：


 1 <= gifts.length <= 10³
 1 <= gifts[i] <= 10⁹
 1 <= k <= 10³


 Related Topics 数组 模拟 堆（优先队列） 👍 57 👎 0

*/

/*
题型: 堆
题目: 从数量最多的堆取走礼物
*/

// leetcode submit region begin(Prohibit modification and deletion)
func pickGifts(gifts []int, k int) int64 {
	q := &pq{}
	for _, gift := range gifts {
		q.Push(gift)
	}
	heap.Init(q)
	for k > 0 {
		x := heap.Pop(q).(int)
		heap.Push(q, int(math.Floor(math.Sqrt(float64(x)))))
		k--
	}

	var res int64
	for q.Len() > 0 {
		res += int64(q.Pop().(int))
	}
	return res
}

type pq []int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *pq) Pop() interface{} {
	n := len(*q)
	res := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
