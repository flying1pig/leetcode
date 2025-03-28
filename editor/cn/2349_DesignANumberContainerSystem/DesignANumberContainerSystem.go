package main

import (
	"container/heap"
	"sort"
)

/**
设计一个数字容器系统，可以实现以下功能：


 在系统中给定下标处 插入 或者 替换 一个数字。
 返回 系统中给定数字的最小下标。


 请你实现一个 NumberContainers 类：


 NumberContainers() 初始化数字容器系统。
 void change(int index, int number) 在下标 index 处填入 number 。如果该下标 index 处已经有数字了，那么
用 number 替换该数字。
 int find(int number) 返回给定数字 number 在系统中的最小下标。如果系统中没有 number ，那么返回 -1 。




 示例：


输入：
["NumberContainers", "find", "change", "change", "change", "change", "find",
"change", "find"]
[[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]
输出：
[null, -1, null, null, null, null, 1, null, 2]

解释：
NumberContainers nc = new NumberContainers();
nc.find(10); // 没有数字 10 ，所以返回 -1 。
nc.change(2, 10); // 容器中下标为 2 处填入数字 10 。
nc.change(1, 10); // 容器中下标为 1 处填入数字 10 。
nc.change(3, 10); // 容器中下标为 3 处填入数字 10 。
nc.change(5, 10); // 容器中下标为 5 处填入数字 10 。
nc.find(10); // 数字 10 所在的下标为 1 ，2 ，3 和 5 。因为最小下标为 1 ，所以返回 1 。
nc.change(1, 20); // 容器中下标为 1 处填入数字 20 。注意，下标 1 处之前为 10 ，现在被替换为 20 。
nc.find(10); // 数字 10 所在下标为 2 ，3 和 5 。最小下标为 2 ，所以返回 2 。




 提示：


 1 <= index, number <= 10⁹
 调用 change 和 find 的 总次数 不超过 10⁵ 次。


 Related Topics 设计 哈希表 有序集合 堆（优先队列） 👍 35 👎 0

*/

/*
题型: 堆
题目: 设计数字容器系统
*/

// leetcode submit region begin(Prohibit modification and deletion)
type NumberContainers struct {
	m  map[int]int
	ms map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*hp{}}
}

func (n NumberContainers) Change(index int, number int) {
	n.m[index] = number
	if n.ms[number] == nil {
		n.ms[number] = &hp{}
	}
	heap.Push(n.ms[number], index) // 直接添加新数据，后面 find 再删除旧的
}

func (n NumberContainers) Find(number int) int {
	h, ok := n.ms[number]
	if !ok {
		return -1
	}
	for h.Len() > 0 && n.m[h.IntSlice[0]] != number {
		heap.Pop(h)
	}
	if h.Len() == 0 {
		return -1
	}
	return h.IntSlice[0]
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
