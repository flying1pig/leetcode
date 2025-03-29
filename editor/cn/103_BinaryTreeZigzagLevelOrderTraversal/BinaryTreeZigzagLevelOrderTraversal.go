package main

/**
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -100 <= Node.val <= 100


 Related Topics 树 广度优先搜索 二叉树 👍 954 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		nxt := []*TreeNode{}
		vals := make([]int, len(cur)) // 大小已知
		for i, node := range cur {
			if len(ans)%2 > 0 {
				vals[len(cur)-1-i] = node.Val // 倒着添加
			} else {
				vals[i] = node.Val
			}
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		ans = append(ans, vals)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
