package main

import (
	"container/heap"
	"sort"
)

/**
现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。

 实现 SmallestInfiniteSet 类：


 SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
 int popSmallest() 移除 并返回该无限集中的最小整数。
 void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。




 示例：


输入
["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest",
"addBack", "popSmallest", "popSmallest", "popSmallest"]
[[], [2], [], [], [], [1], [], [], []]
输出
[null, null, 1, 2, 3, null, 1, 4, 5]

解释
SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
smallestInfiniteSet.addBack(2);    // 2 已经在集合中，所以不做任何变更。
smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 是最小的整数，并将其从集合中移除。
smallestInfiniteSet.popSmallest(); // 返回 2 ，并将其从集合中移除。
smallestInfiniteSet.popSmallest(); // 返回 3 ，并将其从集合中移除。
smallestInfiniteSet.addBack(1);    // 将 1 添加到该集合中。
smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 在上一步中被添加到集合中，
                                   // 且 1 是最小的整数，并将其从集合中移除。
smallestInfiniteSet.popSmallest(); // 返回 4 ，并将其从集合中移除。
smallestInfiniteSet.popSmallest(); // 返回 5 ，并将其从集合中移除。



 提示：


 1 <= num <= 1000
 最多调用 popSmallest 和 addBack 方法 共计 1000 次


 Related Topics 设计 哈希表 有序集合 堆（优先队列） 👍 90 👎 0

*/

/*
题型: 堆
题目: 无限集中的最小数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
type SmallestInfiniteSet struct {
	sort.IntSlice
	exist map[int]struct{}
	limit int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		IntSlice: nil,
		exist:    make(map[int]struct{}),
		limit:    1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.Len() == 0 {
		this.limit++
		return this.limit - 1
	}
	e := heap.Pop(this).(int)
	delete(this.exist, e)
	return e
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.limit {
		return
	}
	if _, ok := this.exist[num]; ok {
		return
	}
	this.exist[num] = struct{}{}
	heap.Push(this, num)
}

func (this *SmallestInfiniteSet) Push(e interface{}) {
	this.IntSlice = append(this.IntSlice, e.(int))
}

func (this *SmallestInfiniteSet) Pop() interface{} {
	e := this.IntSlice[this.Len()-1]
	this.IntSlice = this.IntSlice[:this.Len()-1]
	return e
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
