package main

/**
给定一个已排序的链表的头
 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。



 示例 1：


输入：head = [1,1,2]
输出：[1,2]


 示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]




 提示：


 链表中节点数目在范围 [0, 300] 内
 -100 <= Node.val <= 100
 题目数据保证链表已经按升序 排列


 Related Topics 链表 👍 1195 👎 0

*/

/*
题型: 链表
题目: 删除排序链表中的重复元素
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
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil { // 看看下个节点……
		if cur.Next.Val == cur.Val { // 和我一样，删！
			cur.Next = cur.Next.Next
		} else { // 和我不一样，移动到下个节点
			cur = cur.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
