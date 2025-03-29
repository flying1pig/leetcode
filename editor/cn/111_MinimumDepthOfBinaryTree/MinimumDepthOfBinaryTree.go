package main

import "math"

/**
给定一个二叉树，找出其最小深度。

 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

 说明：叶子节点是指没有子节点的节点。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2


 示例 2：


输入：root = [2,null,3,null,4,null,5,null,6]
输出：5




 提示：


 树中节点数的范围在 [0, 10⁵] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1268 👎 0

*/

/*
题型: 二叉树自顶向下DFS
题目: 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	ans := math.MaxInt
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		cnt++
		if node.Left == nil && node.Right == nil { // node 是叶子
			ans = min(ans, cnt)
			return
		}
		dfs(node.Left, cnt)
		dfs(node.Right, cnt)
	}
	dfs(root, 0)
	if root != nil {
		return ans
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

/*
自顶向下:
func minDepth(root *TreeNode) int {
    ans := math.MaxInt
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, cnt int) {
        if node == nil {
            return
        }
        cnt++
        if node.Left == nil && node.Right == nil { // node 是叶子
            ans = min(ans, cnt)
            return
        }
        dfs(node.Left, cnt)
        dfs(node.Right, cnt)
    }
    dfs(root, 0)
    if root != nil {
        return ans
    }
    return 0
}

*/

/*
自底向上

	func minDepth(root *TreeNode) int {
	    if root == nil {
	        return 0
	    }
	    if root.Right == nil {
	        return minDepth(root.Left) + 1
	    }
	    if root.Left == nil {
	        return minDepth(root.Right) + 1
	    }
	    return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
*/
func main() {

}
