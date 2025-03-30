package main

import "sort"

/**
给你一个二进制矩阵 matrix ，它的大小为 m x n ，你可以将 matrix 中的 列 按任意顺序重新排列。

 请你返回最优方案下将 matrix 重新排列后，全是 1 的子矩阵面积。



 示例 1：




输入：matrix = [[0,0,1],[1,1,1],[1,0,1]]
输出：4
解释：你可以按照上图方式重新排列矩阵的每一列。
最大的全 1 子矩阵是上图中加粗的部分，面积为 4 。


 示例 2：




输入：matrix = [[1,0,1,0,1]]
输出：3
解释：你可以按照上图方式重新排列矩阵的每一列。
最大的全 1 子矩阵是上图中加粗的部分，面积为 3 。


 示例 3：


输入：matrix = [[1,1,0],[1,0,1]]
输出：2
解释：由于你只能整列整列重新排布，所以没有比面积为 2 更大的全 1 子矩形。

 示例 4：


输入：matrix = [[0,0],[0,0]]
输出：0
解释：由于矩阵中没有 1 ，没有任何全 1 的子矩阵，所以面积为 0 。



 提示：


 m == matrix.length
 n == matrix[i].length
 1 <= m * n <= 10⁵
 matrix[i][j] 要么是 0 ，要么是 1 。


 Related Topics 贪心 数组 矩阵 排序 👍 94 👎 0

*/

/*
题型: 贪心
题目: 重新排列后的最大子矩阵
*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	a := make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				a[j]++
			} else {
				a[j] = 0
			}
		}
		b := make([]int, n)
		copy(b, a)
		sort.Ints(b)
		//fmt.Println(b)
		for i := n - 1; i >= 0; i-- {
			ans = max(ans, (n-i)*b[i])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
