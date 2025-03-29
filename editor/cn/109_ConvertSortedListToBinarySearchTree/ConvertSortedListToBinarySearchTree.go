package main

/**
给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。



 示例 1:




输入: head = [-10,-3,0,5,9]
输出: [0,-3,9,-10,null,5]
解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。


 示例 2:


输入: head = []
输出: []




 提示:


 head 中的节点数在[0, 2 * 10⁴] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 939 👎 0

*/

/*
题型: 二叉树+链表
题目: 有序链表转换二叉搜索树
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var h *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	length := 0
	h = head
	for head != nil {
		length++
		head = head.Next
	}
	return buildBST(0, length-1)
}
func buildBST(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) >> 1
	left := buildBST(start, mid-1)
	root := &TreeNode{Val: h.Val}
	h = h.Next
	root.Left = left
	root.Right = buildBST(mid+1, end)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
