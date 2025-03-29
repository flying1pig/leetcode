package main

import "math"

/**
给定一个二叉树，我们在树的节点上安装摄像头。

 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

 计算监控树的所有节点所需的最小摄像头数量。



 示例 1：



 输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。


 示例 2：



 输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。


 提示：


 给定树的节点数的范围是 [1, 1000]。
 每个节点的值都是 0。


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 774 👎 0

*/

/*
题型: 树形 DP
题目: 监控二叉树
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
func dfs(node *TreeNode) (int, int, int) {
	if node == nil {
		return math.MaxInt / 2, 0, 0 // 除 2 防止加法溢出
	}
	lChoose, lByFa, lByChildren := dfs(node.Left)
	rChoose, rByFa, rByChildren := dfs(node.Right)
	choose := min(lChoose, lByFa) + min(rChoose, rByFa) + 1
	byFa := min(lChoose, lByChildren) + min(rChoose, rByChildren)
	byChildren := min(lChoose+rByChildren, lByChildren+rChoose, lChoose+rChoose)
	return choose, byFa, byChildren
}

func minCameraCover(root *TreeNode) int {
	choose, _, byChildren := dfs(root)
	return min(choose, byChildren)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
