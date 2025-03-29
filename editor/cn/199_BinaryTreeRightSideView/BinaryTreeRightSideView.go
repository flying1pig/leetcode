package main

/**
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。



 示例 1：


 输入：root = [1,2,3,null,5,null,4]


 输出：[1,3,4]

 解释：



 示例 2：


 输入：root = [1,2,3,4,null,null,null,5]


 输出：[1,3,4,5]

 解释：



 示例 3：


 输入：root = [1,null,3]


 输出：[1,3]

 示例 4：


 输入：root = []


 输出：[]



 提示:


 二叉树的节点个数的范围是 [0,100]

 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1185 👎 0

*/

/*
题型: 二叉树自顶向下DFS
题目: 二叉树的右视图
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
func rightSideView(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) { // 这个深度首次遇到
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1) // 先递归右子树，保证首次遇到的一定是最右边的节点
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
