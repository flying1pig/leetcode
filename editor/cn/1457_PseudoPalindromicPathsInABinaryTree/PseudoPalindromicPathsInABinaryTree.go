package main

/**
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。

 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。



 示例 1：




输入：root = [2,3,1,3,1,null,1]
输出：2
解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排列
[1,2,1] 。


 示例 2：




输入：root = [2,1,1,1,3,null,null,null,null,null,1]
输出：1
解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。


 示例 3：


输入：root = [9]
输出：1




 提示：


 给定二叉树的节点数目在范围 [1, 10⁵] 内
 1 <= Node.val <= 9


 Related Topics 位运算 树 深度优先搜索 广度优先搜索 二叉树 👍 143 👎 0

*/

/*
题型: 二叉树自顶向下DFS
题目: 二叉树中的伪回文路径
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
func pseudoPalindromicPaths(root *TreeNode) int {
	p := make([]int, 10)
	return dfs(root, p)
}

func dfs(node *TreeNode, p []int) int {
	if node == nil {
		return 0
	}
	p[node.Val] ^= 1 // 修改 node.Val 出现次数的奇偶性
	// 返回前，恢复到递归 node 之前的状态（不做这一步就把 node.Val 算到其它路径中了）
	defer func() { p[node.Val] ^= 1 }()
	if node.Left == node.Right { // node 是叶子节点
		if sum(p) <= 1 {
			return 1
		}
		return 0
	}
	return dfs(node.Left, p) + dfs(node.Right, p)
}

func sum(p []int) (s int) {
	for _, x := range p {
		s += x
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
