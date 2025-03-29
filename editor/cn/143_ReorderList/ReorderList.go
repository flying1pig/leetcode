package main

/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：


L0 → L1 → … → Ln - 1 → Ln


 请将其重新排列后变为：


L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



 示例 1：




输入：head = [1,2,3,4]
输出：[1,4,2,3]

 示例 2：




输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]



 提示：


 链表的长度范围为 [1, 5 * 10⁴]
 1 <= node.val <= 1000


 Related Topics 栈 递归 链表 双指针 👍 1572 👎 0

*/

/*
题型: 链表
题目: 重排链表
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
