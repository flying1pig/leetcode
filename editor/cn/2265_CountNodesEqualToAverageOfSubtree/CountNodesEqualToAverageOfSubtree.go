package main

/**
给你一棵二叉树的根节点 root ，找出并返回满足要求的节点数，要求节点的值等于其 子树 中值的 平均值 。

 注意：


 n 个元素的平均值可以由 n 个元素 求和 然后再除以 n ，并 向下舍入 到最近的整数。
 root 的 子树 由 root 和它的所有后代组成。




 示例 1：
 输入：root = [4,8,5,0,1,null,6]
输出：5
解释：
对值为 4 的节点：子树的平均值 (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4 。
对值为 5 的节点：子树的平均值 (5 + 6) / 2 = 11 / 2 = 5 。
对值为 0 的节点：子树的平均值 0 / 1 = 0 。
对值为 1 的节点：子树的平均值 1 / 1 = 1 。
对值为 6 的节点：子树的平均值 6 / 1 = 6 。


 示例 2：
 输入：root = [1]
输出：1
解释：对值为 1 的节点：子树的平均值 1 / 1 = 1。




 提示：


 树中节点数目在范围 [1, 1000] 内
 0 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 二叉树 👍 47 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 统计值等于子树平均值的节点数
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
func averageOfSubtree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		sum, cnt := node.Val, 1
		if node.Left != nil {
			s, c := dfs(node.Left)
			sum += s
			cnt += c
		}
		if node.Right != nil {
			s, c := dfs(node.Right)
			sum += s
			cnt += c
		}
		if node.Val == sum/cnt {
			ans++
		}
		return sum, cnt
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
