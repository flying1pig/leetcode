package main

import "slices"

/**
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回
这颗 二叉树 。



 示例 1:


输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]


 示例 2:


输入：inorder = [-1], postorder = [-1]
输出：[-1]




 提示:


 1 <= inorder.length <= 3000
 postorder.length == inorder.length
 -3000 <= inorder[i], postorder[i] <= 3000
 inorder 和 postorder 都由 不同 的值组成
 postorder 中每一个值都在 inorder 中
 inorder 保证是树的中序遍历
 postorder 保证是树的后序遍历


 Related Topics 树 数组 哈希表 分治 二叉树 👍 1322 👎 0

*/

/*
题型: 创建二叉树
题目: 从中序与后序遍历序列构造二叉树
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 { // 空节点
		return nil
	}
	leftSize := slices.Index(inorder, postorder[n-1]) // 左子树的大小
	left := buildTree(inorder[:leftSize], postorder[:leftSize])
	right := buildTree(inorder[leftSize+1:], postorder[leftSize:n-1])
	return &TreeNode{postorder[n-1], left, right}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
