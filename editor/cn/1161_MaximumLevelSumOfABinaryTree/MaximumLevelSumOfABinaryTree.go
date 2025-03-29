package main

/**
给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。

 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。



 示例 1：




输入：root = [1,7,0,7,-8,null,null]
输出：2
解释：
第 1 层各元素之和为 1，
第 2 层各元素之和为 7 + 0 = 7，
第 3 层各元素之和为 7 + -8 = -1，
所以我们返回第 2 层的层号，它的层内元素之和最大。


 示例 2：


输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
输出：2




 提示：


 树中的节点数在
 [1, 10⁴]范围内

 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 146 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 最大层内元素和
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
func maxLevelSum(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		q = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs
func maxLevelSum(root *TreeNode) (ans int) {
    sum := []int{}
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if level == len(sum) {
            sum = append(sum, node.Val)
        } else {
            sum[level] += node.Val
        }
        if node.Left != nil {
            dfs(node.Left, level+1)
        }
        if node.Right != nil {
            dfs(node.Right, level+1)
        }
    }
    dfs(root, 0)
    for i, s := range sum {
        if s > sum[ans] {
            ans = i
        }
    }
    return ans + 1 // 层号从 1 开始
}

*/

func main() {

}
