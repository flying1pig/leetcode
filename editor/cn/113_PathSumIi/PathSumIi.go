package main

import "slices"

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

 叶子节点 是指没有子节点的节点。







 示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]


 示例 2：


输入：root = [1,2,3], targetSum = 5
输出：[]


 示例 3：


输入：root = [1,2], targetSum = 0
输出：[]




 提示：


 树中节点总数在范围 [0, 5000] 内
 -1000 <= Node.val <= 1000
 -1000 <= targetSum <= 1000


 Related Topics 树 深度优先搜索 回溯 二叉树 👍 1172 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		left -= node.Val
		// node.Left == node.Right 相当于判断左右节点是否均为 nil
		if node.Left == node.Right && left == 0 {
			ans = append(ans, slices.Clone(path))
		} else {
			dfs(node.Left, left)
			dfs(node.Right, left)
		}
		path = path[:len(path)-1] // 恢复现场
	}

	dfs(root, targetSum)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
