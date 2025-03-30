package main

/**
给你一个整数数组 nums 。每次 move 操作将会选择任意一个满足 0 <= i < nums.length 的下标 i，并将 nums[i] 递增 1。


 返回使 nums 中的每个值都变成唯一的所需要的最少操作次数。

 生成的测试用例保证答案在 32 位整数范围内。







 示例 1：


输入：nums = [1,2,2]
输出：1
解释：经过一次 move 操作，数组将变为 [1, 2, 3]。


 示例 2：


输入：nums = [3,2,1,2,1,7]
输出：6
解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。



提示：


 1 <= nums.length <= 10⁵
 0 <= nums[i] <= 10⁵


 Related Topics 贪心 数组 计数 排序 👍 263 👎 0

*/

/*
题型: 贪心
题目: 使数组唯一的最小增量
*/

// leetcode submit region begin(Prohibit modification and deletion)
const N int = 200005 //全部重复也无妨
func minIncrementForUnique(nums []int) int {
	set := make([]int, N)
	for _, num := range nums {
		set[num]++
	}
	ans := 0
	for i := 0; i < N; i++ {
		if set[i] != 0 {
			for set[i] > 1 {
				j := i
				for ; j < N && set[j] != 0; j++ {
				}
				set[j] = 1
				ans += j - i
				set[i]--
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
