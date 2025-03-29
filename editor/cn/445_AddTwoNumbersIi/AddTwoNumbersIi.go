package main

/**
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

 你可以假设除了数字 0 之外，这两个数字都不会以零开头。



 示例1：




输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]


 示例2：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]


 示例3：


输入：l1 = [0], l2 = [0]
输出：[0]




 提示：


 链表的长度范围为 [1, 100]
 0 <= node.val <= 9
 输入数据保证链表代表的数字无前导 0




 进阶：如果输入链表不能翻转该如何解决？

 Related Topics 栈 链表 数学 👍 758 👎 0

*/

/*
题型: 链表
题目: 两数相加 II
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 视频讲解 https://www.bilibili.com/video/BV1sd4y1x7KN/
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func addTwo(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点
	cur := dummy
	carry := 0                                 // 进位
	for l1 != nil || l2 != nil || carry != 0 { // 有一个不是空节点，或者还有进位，就继续迭代
		if l1 != nil {
			carry += l1.Val // 节点值和进位加在一起
		}
		if l2 != nil {
			carry += l2.Val // 节点值和进位加在一起
		}
		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry /= 10                           // 新的进位
		cur = cur.Next                        // 下一个节点
		if l1 != nil {
			l1 = l1.Next // 下一个节点
		}
		if l2 != nil {
			l2 = l2.Next // 下一个节点
		}
	}
	return dummy.Next // 哨兵节点的下一个节点就是头节点
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2) // l1 和 l2 反转后，就变成【2. 两数相加】了
	l3 := addTwo(l1, l2)
	return reverseList(l3)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
