package main

/**
给你一棵 完美 二叉树的根节点 root ，请你反转这棵树中每个 奇数 层的节点值。


 例如，假设第 3 层的节点值是 [2,1,3,4,7,11,29,18] ，那么反转后它应该变成 [18,29,11,7,4,3,1,2] 。


 反转后，返回树的根节点。

 完美 二叉树需满足：二叉树的所有父节点都有两个子节点，且所有叶子节点都在同一层。

 节点的 层数 等于该节点到根节点之间的边数。



 示例 1：


输入：root = [2,3,5,8,13,21,34]
输出：[2,5,3,8,13,21,34]
解释：
这棵树只有一个奇数层。
在第 1 层的节点分别是 3、5 ，反转后为 5、3 。


 示例 2：


输入：root = [7,13,11]
输出：[7,11,13]
解释：
在第 1 层的节点分别是 13、11 ，反转后为 11、13 。


 示例 3：


输入：root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
输出：[0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
解释：奇数层由非零值组成。
在第 1 层的节点分别是 1、2 ，反转后为 2、1 。
在第 3 层的节点分别是 1、1、1、1、2、2、2、2 ，反转后为 2、2、2、2、1、1、1、1 。




 提示：


 树中的节点数目在范围 [1, 2¹⁴] 内
 0 <= Node.val <= 10⁵
 root 是一棵 完美 二叉树


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 86 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 反转二叉树的奇数层
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
func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	for level := 0; q[0].Left != nil; level ^= 1 {
		next := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			next = append(next, node.Left, node.Right)
		}
		q = next
		if level == 0 {
			for i, n := 0, len(q); i < n/2; i++ {
				x, y := q[i], q[n-1-i]
				x.Val, y.Val = y.Val, x.Val
			}
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs
func dfs(node1 *TreeNode, node2 *TreeNode, isOddLevel bool) {
	if node1 == nil {
		return
	}
	if isOddLevel {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	dfs(node1.Left, node2.Right, !isOddLevel)
	dfs(node1.Right, node2.Left, !isOddLevel)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, true)
	return root
}

*/

func main() {

}
