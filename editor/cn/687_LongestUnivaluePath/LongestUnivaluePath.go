package main

/**
给定一个二叉树的
 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。

 两个节点之间的路径长度 由它们之间的边数表示。



 示例 1:




输入：root = [5,4,5,1,1,5]
输出：2


 示例 2:




输入：root = [1,4,5,4,4,5]
输出：2




 提示:


 树的节点数的范围是
 [0, 10⁴]
 -1000 <= Node.val <= 1000
 树的深度将不超过 1000


 Related Topics 树 深度优先搜索 二叉树 👍 826 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 最长同值路径
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
func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 下面 +1 后，对于叶子节点就刚好是 0
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		if node.Left != nil && node.Left.Val != node.Val {
			lLen = 0 // 链长视作 0
		}
		if node.Right != nil && node.Right.Val != node.Val {
			rLen = 0 // 链长视作 0
		}
		ans = max(ans, lLen+rLen) // 两条链拼成路径
		return max(lLen, rLen)    // 当前子树最大链长
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
