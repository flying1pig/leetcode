package main

/**
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。



 示例 1：


 输入：root = [1,null,2,3]


 输出：[1,2,3]

 解释：



 示例 2：


 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]


 输出：[1,2,4,5,6,7,3,8,9]

 解释：



 示例 3：


 输入：root = []


 输出：[]

 示例 4：


 输入：root = [1]


 输出：[1]



 提示：


 树中节点数目在范围 [0, 100] 内
 -100 <= Node.val <= 100




 进阶：递归算法很简单，你可以通过迭代算法完成吗？

 Related Topics 栈 树 深度优先搜索 二叉树 👍 1321 👎 0

*/

/*
题型: 二叉树遍历
题目: 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
迭代:
func preorderTraversal(root *TreeNode) (vals []int) {
    stack := []*TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            vals = append(vals, node.Val)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    return
}

*/

func main() {

}
