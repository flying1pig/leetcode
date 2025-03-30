package main

import "slices"

/**
给你一个大小为 n x m 的二维矩阵 grid ，以及一个长度为 n 的整数数组 limits ，和一个整数 k 。你的目标是从矩阵 grid 中提取出 至多
 k 个元素，并计算这些元素的最大总和，提取时需满足以下限制：


 从 grid 的第 i 行提取的元素数量不超过 limits[i] 。


 返回最大总和。



 示例 1：


 输入：grid = [[1,2],[3,4]], limits = [1,2], k = 2


 输出：7

 解释：


 从第 2 行提取至多 2 个元素，取出 4 和 3 。
 至多提取 2 个元素时的最大总和 4 + 3 = 7 。


 示例 2：


 输入：grid = [[5,3,7],[8,2,6]], limits = [2,2], k = 3


 输出：21

 解释：


 从第 1 行提取至多 2 个元素，取出 7 。
 从第 2 行提取至多 2 个元素，取出 8 和 6 。
 至多提取 3 个元素时的最大总和 7 + 8 + 6 = 21 。




 提示：


 n == grid.length == limits.length
 m == grid[i].length
 1 <= n, m <= 500
 0 <= grid[i][j] <= 10⁵
 0 <= limits[i] <= m
 0 <= k <= min(n * m, sum(limits))


 Related Topics 贪心 数组 矩阵 排序 堆（优先队列） 👍 4 👎 0

*/

/*
题型: 贪心
题目: 提取至多 K 个元素的最大总和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(grid [][]int, limits []int, k int) (ans int64) {
	a := []int{}
	cmp := func(a, b int) int { return b - a }
	for i, row := range grid {
		slices.SortFunc(row, cmp)
		a = append(a, row[:limits[i]]...)
	}
	slices.SortFunc(a, cmp)
	for _, x := range a[:k] {
		ans += int64(x)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
