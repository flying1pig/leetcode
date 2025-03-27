package main

/**
给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。


 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。




 示例 1:


输入: nums = [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。


 示例 2:


输入: nums = [1,2,3,4,3]
输出: [2,3,4,-1,4]




 提示:


 1 <= nums.length <= 10⁴
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 栈 数组 单调栈 👍 1040 👎 0

*/

/*
题型: 单调栈
题目: 下一个更大元素 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for i := 0; i < 2*n; i++ {
		x := nums[i%n]
		for len(st) > 0 && x > nums[st[len(st)-1]] {
			// x 是 nums[st[len(st)-1]] 的下一个更大元素
			// 既然 nums[st[len(st)-1]] 已经算出答案，则从栈顶弹出
			ans[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
从右到左
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    st := []int{}
    for i := 2*n - 1; i >= 0; i-- {
        x := nums[i%n]
        for len(st) > 0 && x >= st[len(st)-1] {
            // 由于 x 的出现，栈顶元素永远不会是左边元素的「下一个更大元素」
            st = st[:len(st)-1]
        }
        if i < n && len(st) > 0 {
            ans[i] = st[len(st)-1]
        }
        st = append(st, x)
    }
    return ans
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
从左到右
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    st := []int{}
    for i := 0; i < 2 * n; i++ {
        x := nums[i % n]
        for len(st) > 0 && x > nums[st[len(st)-1]] {
            // x 是 nums[st[len(st)-1]] 的下一个更大元素
            // 既然 nums[st[len(st)-1]] 已经算出答案，则从栈顶弹出
            ans[st[len(st)-1]] = x
            st = st[:len(st)-1]
        }
        if i < n {
            st = append(st, i)
        }
    }
    return ans
}
时间复杂度：O(n)
空间复杂度：O(n)
*/

func main() {

}
