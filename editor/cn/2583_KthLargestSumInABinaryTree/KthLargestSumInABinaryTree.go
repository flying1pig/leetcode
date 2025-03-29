package main

import "slices"

/**
给你一棵二叉树的根节点 root 和一个正整数 k 。

 树中的 层和 是指 同一层 上节点值的总和。

 返回树中第 k 大的层和（不一定不同）。如果树少于 k 层，则返回 -1 。

 注意，如果两个节点与根节点的距离相同，则认为它们在同一层。



 示例 1：




输入：root = [5,8,9,2,1,3,7,4,6], k = 2
输出：13
解释：树中每一层的层和分别是：
- Level 1: 5
- Level 2: 8 + 9 = 17
- Level 3: 2 + 1 + 3 + 7 = 13
- Level 4: 4 + 6 = 10
第 2 大的层和等于 13 。


 示例 2：




输入：root = [1,2,null,3], k = 1
输出：3
解释：最大的层和是 3 。




 提示：


 树中的节点数为 n
 2 <= n <= 10⁵
 1 <= Node.val <= 10⁶
 1 <= k <= n


 Related Topics 树 广度优先搜索 二叉树 排序 👍 40 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 二叉树中的第 K 大层和
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
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	a := []int64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := int64(0)
		tmp := q
		q = nil
		for _, node := range tmp {
			sum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		a = append(a, sum)
	}
	n := len(a)
	if k > n {
		return -1
	}
	slices.Sort(a)
	return a[n-k]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
