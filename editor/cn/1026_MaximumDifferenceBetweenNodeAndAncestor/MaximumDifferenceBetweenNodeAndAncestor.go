package main

/**
给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

 （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）



 示例 1：




输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
输出：7
解释：
我们有大量的节点与其祖先的差值，其中一些如下：
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。


 示例 2：


输入：root = [1,null,2,null,0,3]
输出：3




 提示：


 树中的节点数在 2 到 5000 之间。
 0 <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 二叉树 👍 286 👎 0

*/

/*
题型: 二叉树自顶向下DFS
题目: 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, mn, mx int) {
		if node == nil {
			ans = max(ans, mx-mn)
			return
		}
		mn = min(mn, node.Val)
		mx = max(mx, node.Val)
		dfs(node.Left, mn, mx)
		dfs(node.Right, mn, mx)
	}
	dfs(root, root.Val, root.Val)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
自顶向下
func maxAncestorDiff(root *TreeNode) (ans int) {
    var dfs func(*TreeNode, int, int)
    dfs = func(node *TreeNode, mn, mx int) {
        if node == nil {
            ans = max(ans, mx-mn)
            return
        }
        mn = min(mn, node.Val)
        mx = max(mx, node.Val)
        dfs(node.Left, mn, mx)
        dfs(node.Right, mn, mx)
    }
    dfs(root, root.Val, root.Val)
    return
}

*/

/*
自底向上
func maxAncestorDiff(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return math.MaxInt, math.MinInt // 保证空节点不影响 mn 和 mx
        }
        lMn, lMx := dfs(node.Left)
        rMn, rMx := dfs(node.Right)
        mn := min(node.Val, lMn, rMn)
        mx := max(node.Val, lMx, rMx)
        ans = max(ans, node.Val-mn, mx-node.Val)
        return mn, mx
    }
    dfs(root)
    return
}

*/

func main() {

}
