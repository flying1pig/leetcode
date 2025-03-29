package main

/**
给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。如果有多种构造方法，请你返回任意一种。

 如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。



 示例 1：




输入：root = [1,null,2,null,3,null,4,null,null]
输出：[2,1,3,null,null,null,4]
解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。


 示例 2：




输入: root = [2,1,3]
输出: [2,1,3]




 提示：


 树节点的数目在 [1, 10⁴] 范围内。
 1 <= Node.val <= 10⁵


 Related Topics 贪心 树 深度优先搜索 二叉搜索树 分治 二叉树 👍 225 👎 0

*/

/*
题型: 创建二叉树
题目: 将二叉搜索树变平衡
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodes = []int{}

func balanceBST(root *TreeNode) *TreeNode {
	nodes = make([]int, 0)
	inorder(root)
	return build(0, len(nodes)-1)
}

func build(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{
		Val: nodes[mid],
	}
	root.Left = build(start, mid-1)
	root.Right = build(mid+1, end)
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	nodes = append(nodes, root.Val)
	inorder(root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
