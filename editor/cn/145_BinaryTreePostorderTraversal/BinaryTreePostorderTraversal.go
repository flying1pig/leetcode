package main

/**
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



 示例 1：


 输入：root = [1,null,2,3]


 输出：[3,2,1]

 解释：



 示例 2：


 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]


 输出：[4,6,7,5,2,9,8,3,1]

 解释：



 示例 3：


 输入：root = []


 输出：[]

 示例 4：


 输入：root = [1]


 输出：[1]



 提示：


 树中节点的数目在范围 [0, 100] 内
 -100 <= Node.val <= 100




 进阶：递归算法很简单，你可以通过迭代算法完成吗？

 Related Topics 栈 树 深度优先搜索 二叉树 👍 1230 👎 0

*/

/*
题型: 二叉树遍历
题目: 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
迭代:
func postorderTraversal(root *TreeNode) (res []int) {
    stk := []*TreeNode{}
    var prev *TreeNode
    for root != nil || len(stk) > 0 {
        for root != nil {
            stk = append(stk, root)
            root = root.Left
        }
        root = stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        if root.Right == nil || root.Right == prev {
            res = append(res, root.Val)
            prev = root
            root = nil
        } else {
            stk = append(stk, root)
            root = root.Right
        }
    }
    return
}

*/

func main() {

}
