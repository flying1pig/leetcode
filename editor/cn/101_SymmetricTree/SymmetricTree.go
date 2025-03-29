package main

/**
给你一个二叉树的根节点 root ， 检查它是否轴对称。



 示例 1：


输入：root = [1,2,2,3,4,4,3]
输出：true


 示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false




 提示：


 树中节点数目在范围 [1, 1000] 内
 -100 <= Node.val <= 100




 进阶：你可以运用递归和迭代两种方法解决这个问题吗？

 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2920 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 对称二叉树
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
// 在【100. 相同的树】的基础上稍加改动
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
