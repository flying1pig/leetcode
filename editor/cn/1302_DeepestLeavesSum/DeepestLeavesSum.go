package main

/**
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。



 示例 1：




输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15


 示例 2：


输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：19




 提示：


 树中节点数目在范围 [1, 10⁴] 之间。
 1 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 182 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 层数最深叶子节点的和
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
func deepestLeavesSum(root *TreeNode) (sum int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum = 0
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs
func deepestLeavesSum(root *TreeNode) (sum int) {
    maxLevel := -1
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        if level > maxLevel {
            maxLevel = level
            sum = node.Val
        } else if level == maxLevel {
            sum += node.Val
        }
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    dfs(root, 0)
    return
}

*/

func main() {

}
