package main

/**
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。



 每条从根节点到叶节点的路径都代表一个数字：




 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。


 计算从根节点到叶节点生成的 所有数字之和 。

 叶节点 是指没有子节点的节点。



 示例 1：


输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25

 示例 2：


输入：root = [4,9,0,5,1]
输出：1026
解释：
从根到叶子节点路径 4->9->5 代表数字 495
从根到叶子节点路径 4->9->1 代表数字 491
从根到叶子节点路径 4->0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026




 提示：


 树中节点的数目在范围 [1, 1000] 内
 0 <= Node.val <= 9
 树的深度不超过 10


 Related Topics 树 深度优先搜索 二叉树 👍 802 👎 0

*/

/*
题型: 二叉树自顶向下DFS
题目: 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		x = x*10 + node.Val
		if node.Left == node.Right { // node 是叶子节点
			ans += x
			return
		}
		dfs(node.Left, x)
		dfs(node.Right, x)
	}
	dfs(root, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
有返回值的写法:
func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, x int) int {
    if root == nil {
        return 0
    }
    x = x*10 + root.Val
    if root.Left == root.Right { // root 是叶子节点
        return x
    }
    return dfs(root.Left, x) + dfs(root.Right, x)
}

*/

/*
没有返回值的写法:

	func sumNumbers(root *TreeNode) (ans int) {
	    var dfs func(*TreeNode, int)
	    dfs = func(node *TreeNode, x int) {
	        if node == nil {
	            return
	        }
	        x = x*10 + node.Val
	        if node.Left == node.Right { // node 是叶子节点
	            ans += x
	            return
	        }
	        dfs(node.Left, x)
	        dfs(node.Right, x)
	    }
	    dfs(root, 0)
	    return
	}
*/
func main() {

}
