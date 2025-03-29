package main

/**
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



 示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]


 示例 2：


输入：root = []
输出：[]


 示例 3：


输入：root = [1]
输出：[1]




 提示：


 树中节点数目在范围 [0, 100] 内
 -100 <= Node.val <= 100




 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

 Related Topics 栈 树 深度优先搜索 二叉树 👍 2220 👎 0

*/

/*
题型: 二叉树遍历
题目: 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
迭代:
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

*/

func main() {

}
