package main

/**
给定一个二叉树，判断它是否是 平衡二叉树



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true


 示例 2：


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false


 示例 3：


输入：root = []
输出：true




 提示：


 树中的节点数在范围 [0, 5000] 内
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 👍 1590 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 平衡二叉树
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
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := getHeight(node.Left)
	if leftH == -1 {
		return -1 // 提前退出，不再递归
	}
	rightH := getHeight(node.Right)
	if rightH == -1 || abs(leftH-rightH) > 1 {
		return -1
	}
	return max(leftH, rightH) + 1
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
