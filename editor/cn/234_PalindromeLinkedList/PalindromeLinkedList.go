package main

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



 示例 1：


输入：head = [1,2,2,1]
输出：true


 示例 2：


输入：head = [1,2]
输出：false




 提示：


 链表中节点数目在范围[1, 10⁵] 内
 0 <= Node.val <= 9




 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

 Related Topics 栈 递归 链表 双指针 👍 2041 👎 0

*/

/*
题型: 链表
题目: 回文链表
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

func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2 != nil {
		if head.Val != head2.Val { // 不是回文链表
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
