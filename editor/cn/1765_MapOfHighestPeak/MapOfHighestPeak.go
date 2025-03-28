package main

/**
给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。


 如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
 如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。


 你需要按照如下规则给每个单元格安排高度：


 每个格子的高度都必须是非负的。
 如果一个格子是 水域 ，那么它的高度必须为 0 。
 任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）


 找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。

 请你返回一个大小为 m x n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回 任意一个 。




 示例 1：




输入：isWater = [[0,1],[0,0]]
输出：[[1,0],[2,1]]
解释：上图展示了给各个格子安排的高度。
蓝色格子是水域格，绿色格子是陆地格。


 示例 2：




输入：isWater = [[0,0,1],[1,0,0],[0,0,0]]
输出：[[1,1,0],[0,1,1],[1,2,2]]
解释：所有安排方案中，最高可行高度为 2 。
任意安排方案中，只要最高高度为 2 且符合上述规则的，都为可行方案。




 提示：


 m == isWater.length
 n == isWater[i].length
 1 <= m, n <= 1000
 isWater[i][j] 要么是 0 ，要么是 1 。
 至少有 1 个水域格子。

注意：本题与
542 题相同。

 Related Topics 广度优先搜索 数组 矩阵 👍 135 👎 0

*/

/*
题型: 网格bfs
题目: 地图中的最高点
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1 // -1 表示该格子尚未被访问过
		}
	}
	q := []pair{}
	for i, row := range isWater {
		for j, water := range row {
			if water == 1 { // 将所有水域入队
				ans[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && ans[x][y] == -1 {
				ans[x][y] = ans[p.x][p.y] + 1
				q = append(q, pair{x, y})
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
多源广度优先搜索
	题目要求水域的高度必须为 0，因此水域的高度是已经确定的值，我们可以从水域出发，
推导出其余格子的高度：
	首先，计算与水域相邻的格子的高度。对于这些格子来说，其相邻格子中的最小高度即为水域的高度 0，
因此这些格子的高度为 1。
	然后，计算与高度为 1 的格子相邻的、尚未被计算过的格子的高度。对于这些格子来说，
其相邻格子中的最小高度为 1，因此这些格子的高度为 2。
	以此类推，计算出所有格子的高度。
*/

func main() {

}
