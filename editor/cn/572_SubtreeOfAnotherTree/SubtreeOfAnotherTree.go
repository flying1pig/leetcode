package main

/**


 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返
回 false 。



 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。



 示例 1：


输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true


 示例 2：


输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false




 提示：


 root 树上的节点数量范围是 [1, 2000]
 subRoot 树上的节点数量范围是 [1, 1000]
 -10⁴ <= root.val <= 10⁴
 -10⁴ <= subRoot.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 1126 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 另一棵树的子树
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
// 100. 相同的树
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 必须都是 nil
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) ||
		isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
