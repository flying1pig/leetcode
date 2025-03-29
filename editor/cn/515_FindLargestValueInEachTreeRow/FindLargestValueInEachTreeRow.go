package main

import "math"

/**
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。



 示例1：




输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]


 示例2：


输入: root = [1,2,3]
输出: [1,3]




 提示：


 二叉树的节点个数的范围是 [0,10⁴]

 -2³¹ <= Node.val <= 2³¹ - 1




 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 394 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 在每个树行中找最大值
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
func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		maxVal := math.MinInt32
		tmp := q
		q = nil
		for _, node := range tmp {
			maxVal = max(maxVal, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, maxVal)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs
func largestValues(root *TreeNode) (ans []int) {
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, curHeight int) {
        if node == nil {
            return
        }
        if curHeight == len(ans) {
            ans = append(ans, node.Val)
        } else {
            ans[curHeight] = max(ans[curHeight], node.Val)
        }
        dfs(node.Left, curHeight+1)
        dfs(node.Right, curHeight+1)
    }
    dfs(root, 0)
    return
}

*/

func main() {

}
