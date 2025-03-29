package main

/**
给你二叉树的根节点 root 和一个整数 limit ，请你同时删除树中所有 不足节点 ，并返回最终二叉树的根节点。

 假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为 不足节点 ，需要被删除。

 叶子节点，就是没有子节点的节点。



 示例 1：


输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
输出：[1,2,3,4,null,null,7,8,9,null,14]


 示例 2：


输入：root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
输出：[5,4,8,11,null,17,4,7,null,null,null,5]


 示例 3：


输入：root = [1,2,-3,-5,null,4,null], limit = -1
输出：[1,null,-3,4]




 提示：


 树中节点数目在范围 [1, 5000] 内
 -10⁵ <= Node.val <= 10⁵
 -10⁹ <= limit <= 10⁹




 Related Topics 树 深度优先搜索 二叉树 👍 189 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 根到叶路径上的不足节点
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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val
	if root.Left == root.Right { // root 是叶子
		if limit > 0 { // 从根到叶子的路径和小于 limit，删除叶子
			return nil
		}
		return root // 否则不删除
	}
	root.Left = sufficientSubset(root.Left, limit)
	root.Right = sufficientSubset(root.Right, limit)
	if root.Left == nil && root.Right == nil { // 如果儿子都被删除，就删 root
		return nil
	}
	return root // 否则不删 root
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
