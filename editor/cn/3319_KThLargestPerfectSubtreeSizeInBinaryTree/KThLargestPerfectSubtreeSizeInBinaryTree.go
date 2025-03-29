package main

/**
给你一棵 二叉树 的根节点 root 和一个整数k。

 返回第 k 大的 完美二叉子树 的大小，如果不存在则返回 -1。

 完美二叉树 是指所有叶子节点都在同一层级的树，且每个父节点恰有两个子节点。



 示例 1：


 输入： root = [5,3,6,5,2,5,7,1,8,null,null,6,8], k = 2


 输出： 3

 解释：



 完美二叉子树的根节点在图中以黑色突出显示。它们的大小按非递增顺序排列为 [3, 3, 1, 1, 1, 1, 1, 1]。 第 2 大的完美二叉子树的大小是
3。

 示例 2：


 输入： root = [1,2,3,4,5,6,7], k = 1


 输出： 7

 解释：



 完美二叉子树的大小按非递增顺序排列为 [7, 3, 3, 1, 1, 1, 1]。最大的完美二叉子树的大小是 7。

 示例 3：


 输入： root = [1,2,3,null,4], k = 3


 输出： -1

 解释：



 完美二叉子树的大小按非递增顺序排列为 [1, 1]。完美二叉子树的数量少于 3。



 提示：


 树中的节点数目在 [1, 2000] 范围内。
 1 <= Node.val <= 2000
 1 <= k <= 1024


 Related Topics 树 深度优先搜索 二叉树 排序 👍 2 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 第 K 大的完美二叉子树的大小
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
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	cnt := [10]int{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH == -2 || leftH != rightH {
			return -2 // 不合法
		}
		cnt[leftH+1]++
		return leftH + 1
	}
	dfs(root)

	for i := len(cnt) - 1; i >= 0; i-- {
		c := cnt[i]
		if c >= k {
			return 1<<(i+1) - 1
		}
		k -= c
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
