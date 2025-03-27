package main

/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

 求在该柱状图中，能够勾勒出来的矩形的最大面积。



 示例 1:




输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10


 示例 2：




输入： heights = [2,4]
输出： 4



 提示：


 1 <= heights.length <=10⁵
 0 <= heights[i] <= 10⁴


 Related Topics 栈 数组 单调栈 👍 2897 👎 0

*/

/*
题型: 单调栈
题目: 柱状图中最大的矩形
*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left := make([]int, n)
	st := []int{}
	for i, x := range heights {
		for len(st) > 0 && x <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i, h := range heights {
		ans = max(ans, h*(right[i]-left[i]-1))
	}
	return ans
}

//时间复杂度: o(n)
//空间复杂度: o(n)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
