package main

/**
给定二叉树的根节点 root ，返回所有左叶子之和。



 示例 1：




输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24


 示例 2:


输入: root = [1]
输出: 0




 提示:


 节点数在 [1, 1000] 范围内
 -1000 <= Node.val <= 1000




 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 766 👎 0

*/

/*
题型: 二叉树遍历
题目: 左叶子之和
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
func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs:
func isLeafNode(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
    if node.Left != nil {
        if isLeafNode(node.Left) {
            ans += node.Left.Val
        } else {
            ans += dfs(node.Left)
        }
    }
    if node.Right != nil && !isLeafNode(node.Right) {
        ans += dfs(node.Right)
    }
    return
}

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root)
}

*/

/*
bfs:

	func isLeafNode(node *TreeNode) bool {
	    return node.Left == nil && node.Right == nil
	}

	func sumOfLeftLeaves(root *TreeNode) (ans int) {
	    if root == nil {
	        return
	    }
	    q := []*TreeNode{root}
	    for len(q) > 0 {
	        node := q[0]
	        q = q[1:]
	        if node.Left != nil {
	            if isLeafNode(node.Left) {
	                ans += node.Left.Val
	            } else {
	                q = append(q, node.Left)
	            }
	        }
	        if node.Right != nil && !isLeafNode(node.Right) {
	            q = append(q, node.Right)
	        }
	    }
	    return
	}
*/
func main() {

}
