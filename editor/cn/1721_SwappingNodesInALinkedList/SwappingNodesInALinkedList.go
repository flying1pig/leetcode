package main

/**
给你链表的头节点 head 和一个整数 k 。

 交换 链表正数第 k 个节点和倒数第 k 个节点的值后，返回链表的头节点（链表 从 1 开始索引）。



 示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[1,4,3,2,5]


 示例 2：


输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
输出：[7,9,6,6,8,7,3,0,9,5]


 示例 3：


输入：head = [1], k = 1
输出：[1]


 示例 4：


输入：head = [1,2], k = 1
输出：[2,1]


 示例 5：


输入：head = [1,2,3], k = 2
输出：[1,2,3]




 提示：


 链表中节点的数目是 n
 1 <= k <= n <= 10⁵
 0 <= Node.val <= 100


 Related Topics 链表 双指针 👍 112 👎 0

*/

/*
题型: 链表
题目: 交换链表中的节点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	ans := &ListNode{Next: head}
	temp1, temp2, temp3 := ans, ans, ans
	//遍历k次后 temp1为head第k个节点
	for i := 0; i < k; i++ {
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	//当temp1为nil时 temp3 为head倒数第k个节点
	for temp1 != nil {
		temp3 = temp3.Next
		temp1 = temp1.Next
	}
	//交换
	numberk := temp2.Val
	temp2.Val = temp3.Val
	temp3.Val = numberk
	return ans.Next

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
