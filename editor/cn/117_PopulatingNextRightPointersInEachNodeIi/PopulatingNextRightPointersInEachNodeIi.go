package main

/**
给定一个二叉树：


struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}

 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

 初始状态下，所有 next 指针都被设置为 NULL 。



 示例 1：


输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连
接），'#' 表示每层的末尾。

 示例 2：


输入：root = []
输出：[]




 提示：


 树中的节点数在范围 [0, 6000] 内
 -100 <= Node.val <= 100


 进阶：


 你只能使用常量级额外空间。
 使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。





 Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 911 👎 0

*/

/*
题型: 二叉树+链表
题目: 填充每个节点的下一个右侧节点指针 II
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	pre := []*Node{}
	var dfs func(*Node, int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth == len(pre) { // node 是这一层最左边的节点
			pre = append(pre, node)
		} else { // pre[depth] 是 node 左边的节点
			pre[depth].Next = node // node 左边的节点指向 node
			pre[depth] = node
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0) // 根节点的深度为 0
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

/*
bfs
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    q := []*Node{root}
    for q != nil {
        tmp := q
        q = nil
        for i, node := range tmp {
            if i > 0 { // 连接同一层的两个相邻节点
                tmp[i-1].Next = node
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
    return root
}

*/

func main() {

}
