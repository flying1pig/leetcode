package main

import "slices"

/**
给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的
后序遍历，重构并返回二叉树。

 如果存在多个答案，您可以返回其中 任何 一个。



 示例 1：




输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]


 示例 2:


输入: preorder = [1], postorder = [1]
输出: [1]




 提示：


 1 <= preorder.length <= 30
 1 <= preorder[i] <= preorder.length
 preorder 中所有值都 不同
 postorder.length == preorder.length
 1 <= postorder[i] <= postorder.length
 postorder 中所有值都 不同
 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历


 Related Topics 树 数组 哈希表 分治 二叉树 👍 409 👎 0

*/

/*
题型: 创建二叉树
题目: 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { // 空节点
		return nil
	}
	if n == 1 { // 叶子节点
		return &TreeNode{Val: preorder[0]}
	}
	leftSize := slices.Index(postorder, preorder[1]) + 1 // 左子树的大小
	left := constructFromPrePost(preorder[1:1+leftSize], postorder[:leftSize])
	right := constructFromPrePost(preorder[1+leftSize:], postorder[leftSize:n-1])
	return &TreeNode{preorder[0], left, right}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
