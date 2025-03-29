package main

/**
给你一棵二叉树的根节点，返回该树的 直径 。

 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

 两节点之间路径的 长度 由它们之间边数表示。



 示例 1：


输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。


 示例 2：


输入：root = [1,2]
输出：1




 提示：


 树中节点数目在范围 [1, 10⁴] 内
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 二叉树 👍 1695 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
