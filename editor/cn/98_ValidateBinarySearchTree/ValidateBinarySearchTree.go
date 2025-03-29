package main

import "math"

/**
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

 有效 二叉搜索树定义如下：


 节点的左子树只包含 小于 当前节点的数。
 节点的右子树只包含 大于 当前节点的数。
 所有左子树和右子树自身必须也是二叉搜索树。




 示例 1：


输入：root = [2,1,3]
输出：true


 示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。




 提示：


 树中节点数目范围在[1, 10⁴] 内
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 2550 👎 0

*/

/*
题型: 二叉搜索树
题目: 验证二叉搜索树
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
func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	x := node.Val
	return left < x && x < right &&
		dfs(node.Left, left, x) &&
		dfs(node.Right, x, right)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
前序遍历
func dfs(node *TreeNode, left, right int) bool {
    if node == nil {
        return true
    }
    x := node.Val
    return left < x && x < right &&
        dfs(node.Left, left, x) &&
        dfs(node.Right, x, right)
}

func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt, math.MaxInt)
}
*/

/*
中序遍历
func isValidBST(root *TreeNode) bool {
    pre := math.MinInt
    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return true
        }
        if !dfs(node.Left) || node.Val <= pre {
            return false
        }
        pre = node.Val
        return dfs(node.Right)
    }
    return dfs(root)
}
*/

/*
后续遍历

	func dfs(node *TreeNode) (int, int) {
	    if node == nil {
	        return math.MaxInt, math.MinInt
	    }
	    lMin, lMax := dfs(node.Left)
	    rMin, rMax := dfs(node.Right)
	    x := node.Val
	    if x <= lMax || x >= rMin { // 也可以在递归完左子树之后立刻判断，如果发现不是二叉搜索树，就不用递归右子树了
	        return math.MinInt, math.MaxInt
	    }
	    return min(lMin, x), max(rMax, x)
	}

	func isValidBST(root *TreeNode) bool {
	    _, mx := dfs(root)
	    return mx != math.MaxInt
	}
*/
func main() {

}
