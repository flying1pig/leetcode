package main

/**
在给定的 m x n 网格
 grid 中，每个单元格可以有以下三个值之一：


 值 0 代表空单元格；
 值 1 代表新鲜橘子；
 值 2 代表腐烂的橘子。


 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。



 示例 1：




输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4


 示例 2：


输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。


 示例 3：


输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 10
 grid[i][j] 仅为 0、1 或 2


 Related Topics 广度优先搜索 数组 矩阵 👍 998 👎 0

*/

/*
题型: 网格bfs
题目: 腐烂的橘子
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 四方向

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	q := []pair{}
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				fresh++ // 统计新鲜橘子个数
			} else if x == 2 {
				q = append(q, pair{i, j}) // 一开始就腐烂的橘子
			}
		}
	}

	ans := 0
	for fresh > 0 && len(q) > 0 {
		ans++ // 经过一分钟
		tmp := q
		q = []pair{}
		for _, p := range tmp { // 已经腐烂的橘子
			for _, d := range directions { // 四方向
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 { // 新鲜橘子
					fresh--
					grid[i][j] = 2 // 变成腐烂橘子
					q = append(q, pair{i, j})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
