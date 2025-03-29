package main

import "strconv"

/**
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

 叶子节点 是指没有子节点的节点。

 示例 1：


输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]


 示例 2：


输入：root = [1]
输出：["1"]




 提示：


 树中节点的数目在范围 [1, 100] 内
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1219 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) (ans []string) {
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == node.Right { // 叶子节点
			ans = append(ans, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
