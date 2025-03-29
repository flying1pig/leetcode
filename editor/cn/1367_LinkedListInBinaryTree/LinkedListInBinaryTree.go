package main

/**
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。


 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。



 示例 1：



 输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1
,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。


 示例 2：



 输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
null,1,3]
输出：true


 示例 3：

 输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。




 提示：


 二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
 链表包含的节点数目在 1 到 100 之间。
 二叉树包含的节点数目在 1 到 2500 之间。


 Related Topics 树 深度优先搜索 链表 二叉树 👍 224 👎 0

*/

/*
题型: 二叉树+链表
题目: 二叉树中的链表
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	// mod 和 base 随机其中一个就行，无需两个都随机
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack

	n := 0        // 链表长度
	powBase := 1  // base^(n-1)
	listHash := 0 // 多项式哈希 s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	for ; head != nil; head = head.Next {
		n++
		if n > 1 {
			powBase = powBase * base % mod
		}
		listHash = (listHash*base + head.Val) % mod // 秦九韶算法计算多项式哈希
	}

	st := []int{}
	var dfs func(*TreeNode, int) bool
	dfs = func(t *TreeNode, hash int) bool {
		if t == nil { // 无法继续匹配
			return false
		}
		st = append(st, t.Val)
		hash = (hash*base + t.Val) % mod // 移入窗口
		if len(st) >= n {
			if hash == listHash {
				return true
			}
			hash = (hash - powBase*st[len(st)-n]%mod + mod) % mod // 移出窗口
		}
		defer func() { st = st[:len(st)-1] }() // 恢复现场
		return dfs(t.Left, hash) || dfs(t.Right, hash)
	}
	return dfs(root, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
