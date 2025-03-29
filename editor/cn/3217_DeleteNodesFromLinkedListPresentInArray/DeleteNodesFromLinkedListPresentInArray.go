package main

/**
给你一个整数数组 nums 和一个链表的头节点 head。从链表中移除所有存在于 nums 中的节点后，返回修改后的链表的头节点。



 示例 1：


 输入： nums = [1,2,3], head = [1,2,3,4,5]


 输出： [4,5]

 解释：



 移除数值为 1, 2 和 3 的节点。

 示例 2：


 输入： nums = [1], head = [1,2,1,2,1,2]


 输出： [2,2,2]

 解释：



 移除数值为 1 的节点。

 示例 3：


 输入： nums = [5], head = [1,2,3,4]


 输出： [1,2,3,4]

 解释：



 链表中不存在值为 5 的节点。



 提示：


 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁵
 nums 中的所有元素都是唯一的。
 链表中的节点数在 [1, 10⁵] 的范围内。
 1 <= Node.val <= 10⁵
 输入保证链表中至少有一个值没有在 nums 中出现过。


 Related Topics 数组 哈希表 链表 👍 12 👎 0

*/

/*
题型: 链表
题目: 从链表中移除在数组中存在的节点
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	has := make(map[int]bool, len(nums)) // 预分配空间
	for _, x := range nums {
		has[x] = true
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if has[cur.Next.Val] {
			cur.Next = cur.Next.Next // 删除
		} else {
			cur = cur.Next // 向后移动
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
