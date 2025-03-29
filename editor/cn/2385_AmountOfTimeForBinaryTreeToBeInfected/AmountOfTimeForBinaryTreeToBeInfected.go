package main

/**
给你一棵二叉树的根节点 root ，二叉树中节点的值 互不相同 。另给你一个整数 start 。在第 0 分钟，感染 将会从值为 start 的节点开始爆发。


 每分钟，如果节点满足以下全部条件，就会被感染：


 节点此前还没有感染。
 节点与一个已感染节点相邻。


 返回感染整棵树需要的分钟数。



 示例 1：
 输入：root = [1,5,3,null,4,10,6,9,2], start = 3
输出：4
解释：节点按以下过程被感染：
- 第 0 分钟：节点 3
- 第 1 分钟：节点 1、10、6
- 第 2 分钟：节点5
- 第 3 分钟：节点 4
- 第 4 分钟：节点 9 和 2
感染整棵树需要 4 分钟，所以返回 4 。


 示例 2：
 输入：root = [1], start = 1
输出：0
解释：第 0 分钟，树中唯一一个节点处于感染状态，返回 0 。




 提示：


 树中节点的数目在范围 [1, 10⁵] 内
 1 <= Node.val <= 10⁵
 每个节点的值 互不相同
 树中必定存在值为 start 的节点


 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 115 👎 0

*/

/*
题型: 二叉树自底向上DFS
题目: 感染二叉树需要的总时间
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
func amountOfTime(root *TreeNode, start int) (ans int) {
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		lLen, lFound := dfs(node.Left)
		rLen, rFound := dfs(node.Right)
		if node.Val == start {
			// 计算子树 start 的最大深度
			// 注意这里和方法一的区别，max 后面没有 +1，所以算出的也是最大深度
			ans = max(lLen, rLen)
			return 1, true // 找到了 start
		}
		if lFound || rFound {
			// 只有在左子树或右子树包含 start 时，才能更新答案
			ans = max(ans, lLen+rLen) // 两条链拼成直径
			// 保证 start 是直径端点
			if lFound {
				return lLen + 1, true
			}
			return rLen + 1, true
		}
		return max(lLen, rLen) + 1, false
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
递归时，除了返回当前子树的最大链长加一，还需要返回一个布尔值，表示当前子树是否包含 start。
如果当前节点是空节点，返回 0 和 false。
设左子树的返回的链长为 lLen，右子树返回的链长为 rLen。
如果当前节点值等于 start，初始化答案为 max(lLen,rLen)，即子树 start 的最大深度。然后返回 1 和 true。
如果左右子树都不包含 start，返回 max(lLen,rLen)+1。
如果左子树或右子树包含 start，像计算直径那样，用 lLen+rLen 更新答案的最大值。如果左子树包含 start，则返回 lLen 和 true，否则返回 rLen 和 true。这种返回方式可以保证 lLen+rLen 一定是端点为 start 的直径长度。

*/

func main() {

}
