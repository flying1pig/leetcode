package main

import "math"

/**
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过
根节点。

 路径和 是路径中各节点值的总和。

 给你一个二叉树的根节点 root ，返回其 最大路径和 。



 示例 1：


输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

 示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42




 提示：


 树中节点数目范围是 [1, 3 * 10⁴]
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 2376 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0 // 没有节点，和为 0
		}
		lVal := dfs(node.Left)                  // 左子树最大链和
		rVal := dfs(node.Right)                 // 右子树最大链和
		ans = max(ans, lVal+rVal+node.Val)      // 两条链拼成路径
		return max(max(lVal, rVal)+node.Val, 0) // 当前子树最大链和（注意这里和 0 取最大值了）
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
