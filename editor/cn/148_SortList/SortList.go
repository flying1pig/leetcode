package main

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。






 示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]


 示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]


 示例 3：


输入：head = []
输出：[]




 提示：


 链表中节点的数目在范围 [0, 5 * 10⁴] 内
 -10⁵ <= Node.val <= 10⁵




 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

 Related Topics 链表 双指针 分治 排序 归并排序 👍 2484 👎 0

*/

/*
题型: 链表
题目: 排序链表
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 876. 链表的中间结点（快慢指针）
func middleNode(head *ListNode) *ListNode {
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow // 记录 slow 的前一个节点
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 断开 slow 的前一个节点和 slow 的连接
	return slow
}

// 21. 合并两个有序链表（双指针）
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := ListNode{} // 用哨兵节点简化代码逻辑
	cur := &dummy       // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	// 如果链表为空或者只有一个节点，无需排序
	if head == nil || head.Next == nil {
		return head
	}
	// 找到中间节点 head2，并断开 head2 与其前一个节点的连接
	// 比如 head=[4,2,1,3]，那么 middleNode 调用结束后 head=[4,2] head2=[1,3]
	head2 := middleNode(head)
	// 分治
	head = sortList(head)
	head2 = sortList(head2)
	// 合并
	return mergeTwoLists(head, head2)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
