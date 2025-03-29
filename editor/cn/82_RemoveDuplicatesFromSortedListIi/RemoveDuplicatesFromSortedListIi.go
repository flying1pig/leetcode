package main

/**
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。



 示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]


 示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]




 提示：


 链表中节点数目在范围 [0, 300] 内
 -100 <= Node.val <= 100
 题目数据保证链表已经按升序 排列


 Related Topics 链表 双指针 👍 1346 👎 0

*/

/*
题型: 链表
题目: 删除排序链表中的重复元素 II
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if cur.Next.Next.Val == val { // 后两个节点值相同
			// 值等于 val 的节点全部删除
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
