package main

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节
点，返回 反转后的链表 。



 示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]


 示例 2：


输入：head = [5], left = 1, right = 1
输出：[5]




 提示：


 链表中节点数目为 n
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n




 进阶： 你可以使用一趟扫描完成反转吗？

 Related Topics 链表 👍 1941 👎 0

*/

/*
题型: 链表
题目: 反转链表 II
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
		pre = cur
		cur = nxt
	}

	// 见视频
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
