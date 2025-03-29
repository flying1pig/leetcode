package main

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



 示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


 示例 2：


输入：head = [1], n = 1
输出：[]


 示例 3：


输入：head = [1,2], n = 1
输出：[1]




 提示：


 链表中结点的数目为 sz
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz




 进阶：你能尝试使用一趟扫描实现吗？

 Related Topics 链表 双指针 👍 3078 👎 0

*/

/*
题型: 链表
题目: 删除链表的倒数第 N 个结点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 由于可能会删除链表头部，用哨兵节点简化代码
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy
	for ; n > 0; n-- {
		right = right.Next // 右指针先向右走 n 步
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next // 左右指针一起走
	}
	left.Next = left.Next.Next // 左指针的下一个节点就是倒数第 n 个节点
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
