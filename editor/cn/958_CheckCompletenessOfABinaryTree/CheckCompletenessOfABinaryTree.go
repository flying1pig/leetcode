package main

/**
给你一棵二叉树的根节点
 root ，请你判断这棵树是否是一棵 完全二叉树 。

 在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含
 1 到
 2ʰ 个节点。



 示例 1：




输入：root = [1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6}）尽可能靠左。


 示例 2：




输入：root = [1,2,3,4,5,null,7]
输出：false
解释：值为 7 的节点不满足条件「节点尽可能靠左」。




 提示：


 树中节点数目在范围 [1, 100] 内
 1 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 308 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	//标记层序遍历的过程中是否有遇到空节点
	empty := false
	//辅助层序遍历的队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		//取一个节点
		node := queue[0]
		//出队列
		queue = queue[1:]
		if node == nil {
			//遇到空节点，把标记改成true
			empty = true
		} else {
			//此时是遍历非空节点，在非空节点之前出现了空节点，就说明并不是满二叉树
			if empty == true {
				return false
			}
			//添加左右子节点，无论是否为空
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
