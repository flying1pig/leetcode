package main

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。



 示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]


 示例 2：


输入：head = []
输出：[]


 示例 3：


输入：head = [1]
输出：[1]




 提示：


 链表中节点的数目在范围 [0, 100] 内
 0 <= Node.val <= 100


 Related Topics 递归 链表 👍 2385 👎 0

*/

/*
题型: 链表
题目: 两两交换链表中的节点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 用哨兵节点简化代码逻辑
	node0 := dummy
	node1 := head
	for node1 != nil && node1.Next != nil { // 至少有两个节点
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2 // 0 -> 2
		node2.Next = node1 // 2 -> 1
		node1.Next = node3 // 1 -> 3

		node0 = node1 // 下一轮交换，0 是 1
		node1 = node3 // 下一轮交换，1 是 3
	}
	return dummy.Next // 返回新链表的头节点
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
