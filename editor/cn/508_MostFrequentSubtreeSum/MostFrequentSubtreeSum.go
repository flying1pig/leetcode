package main

/**
给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。



 示例 1：




输入: root = [5,2,-3]
输出: [2,-3,4]


 示例 2：




输入: root = [5,2,-5]
输出: [2]




 提示:


 节点数在 [1, 10⁴] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 哈希表 二叉树 👍 255 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 出现次数最多的子树元素和
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
func findFrequentTreeSum(root *TreeNode) (ans []int) {
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		if cnt[sum] > maxCnt {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)

	for s, c := range cnt {
		if c == maxCnt {
			ans = append(ans, s)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
