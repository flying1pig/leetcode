package main

/**
给定一个长度为 n 的链表 head

 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。

 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设
置 answer[i] = 0 。



 示例 1：




输入：head = [2,1,5]
输出：[5,5,0]


 示例 2：




输入：head = [2,7,4,3,5]
输出：[7,0,5,5,0]




 提示：


 链表中节点数为 n
 1 <= n <= 10⁴
 1 <= Node.val <= 10⁹


 Related Topics 栈 数组 链表 单调栈 👍 352 👎 0

*/

/*
题型: 链表
题目: 链表中的下一个更大节点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 206. 反转链表
func reverseList(head *ListNode) (pre *ListNode, n int) {
	for cur := head; cur != nil; n++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return
}

func nextLargerNodes(head *ListNode) []int {
	head, n := reverseList(head)
	ans := make([]int, n)
	st := []int{} // 单调栈（节点值）
	for cur := head; cur != nil; cur = cur.Next {
		for len(st) > 0 && st[len(st)-1] <= cur.Val {
			st = st[:len(st)-1] // 弹出无用数据
		}
		n--
		if len(st) > 0 {
			ans[n] = st[len(st)-1]
		}
		st = append(st, cur.Val)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
