package main

import "slices"

/**
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回
其根节点。



 示例 1:


输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]


 示例 2:


输入: preorder = [-1], inorder = [-1]
输出: [-1]




 提示:


 1 <= preorder.length <= 3000
 inorder.length == preorder.length
 -3000 <= preorder[i], inorder[i] <= 3000
 preorder 和 inorder 均 无重复 元素
 inorder 均出现在 preorder
 preorder 保证 为二叉树的前序遍历序列
 inorder 保证 为二叉树的中序遍历序列


 Related Topics 树 数组 哈希表 分治 二叉树 👍 2506 👎 0

*/

/*
题型: 创建二叉树
题目: 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { // 空节点
		return nil
	}
	leftSize := slices.Index(inorder, preorder[0]) // 左子树的大小
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
	return &TreeNode{preorder[0], left, right}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
