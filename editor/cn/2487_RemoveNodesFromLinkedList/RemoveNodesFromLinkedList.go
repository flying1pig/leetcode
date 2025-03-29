package main

/**
给你一个链表的头节点 head 。

 移除每个右侧有一个更大数值的节点。

 返回修改后链表的头节点 head 。



 示例 1：




输入：head = [5,2,13,3,8]
输出：[13,8]
解释：需要移除的节点是 5 ，2 和 3 。
- 节点 13 在节点 5 右侧。
- 节点 13 在节点 2 右侧。
- 节点 8 在节点 3 右侧。


 示例 2：


输入：head = [1,1,1,1]
输出：[1,1,1,1]
解释：每个节点的值都是 1 ，所以没有需要移除的节点。




 提示：


 给定列表中的节点数目在范围 [1, 10⁵] 内
 1 <= Node.val <= 10⁵


 Related Topics 栈 递归 链表 单调栈 👍 126 👎 0

*/

/*
题型: 链表
题目: 从链表中移除节点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	cur := head
	for cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return reverseList(head)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
递归:
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	node := removeNodes(head.Next)
	if node.Val > head.Val { // 返回的链表头一定是最大的
		return node // 删除 head
	}
	head.Next = node // 不删除 head
	return head
}

*/

func main() {

}
