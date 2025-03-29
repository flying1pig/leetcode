package main

/**
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

 假设二叉树中至少有一个节点。



 示例 1:




输入: root = [2,1,3]
输出: 1


 示例 2:




输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7




 提示:


 二叉树的节点个数的范围是 [1,10⁴]

 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 627 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
