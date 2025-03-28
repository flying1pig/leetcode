package main

import (
	"container/heap"
	"sort"
)

/**
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

 请实现 KthLargest 类：


 KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
 int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。




 示例 1：


 输入： ["KthLargest", "add", "add", "add", "add", "add"] [[3, [4, 5, 8, 2]], [3],
[5], [10], [9], [4]]


 输出：[null, 4, 5, 5, 8, 8]

 解释：

 KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]); kthLargest.add(3); //
返回 4 kthLargest.add(5); // 返回 5 kthLargest.add(10); // 返回 5 kthLargest.add(9); /
/ 返回 8 kthLargest.add(4); // 返回 8



 示例 2：


 输入： ["KthLargest", "add", "add", "add", "add"] [[4, [7, 7, 7, 7, 8, 3]], [2], [
10], [9], [9]]


 输出：[null, 7, 7, 7, 8]

 解释： KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
 kthLargest.add(2); // 返回 7
 kthLargest.add(10); // 返回 7
 kthLargest.add(9); // 返回 7
 kthLargest.add(9); // 返回 8


提示：


 0 <= nums.length <= 10⁴
 1 <= k <= nums.length + 1
 -10⁴ <= nums[i] <= 10⁴
 -10⁴ <= val <= 10⁴
 最多调用 add 方法 10⁴ 次


 Related Topics 树 设计 二叉搜索树 二叉树 数据流 堆（优先队列） 👍 493 👎 0

*/

/*
题型: 堆
题目: 数据流中的第 K 大元素
*/

// leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
