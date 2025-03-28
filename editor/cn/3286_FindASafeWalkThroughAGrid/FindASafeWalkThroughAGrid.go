package main

import "math"

/**
给你一个 m x n 的二进制矩形 grid 和一个整数 health 表示你的健康值。

 你开始于矩形的左上角 (0, 0) ，你的目标是矩形的右下角 (m - 1, n - 1) 。

 你可以在矩形中往上下左右相邻格子移动，但前提是你的健康值始终是 正数 。

 对于格子 (i, j) ，如果 grid[i][j] = 1 ，那么这个格子视为 不安全 的，会使你的健康值减少 1 。

 如果你可以到达最终的格子，请你返回 true ，否则返回 false 。

 注意 ，当你在最终格子的时候，你的健康值也必须为 正数 。



 示例 1：


 输入：grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1


 输出：true

 解释：

 沿着下图中灰色格子走，可以安全到达最终的格子。


 示例 2：


 输入：grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3



 输出：false

 解释：

 健康值最少为 4 才能安全到达最后的格子。


 示例 3：


 输入：grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5


 输出：true

 解释：

 沿着下图中灰色格子走，可以安全到达最终的格子。



 任何不经过格子 (1, 1) 的路径都是不安全的，因为你的健康值到达最终格子时，都会小于等于 0 。



 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 50
 2 <= m * n
 1 <= health <= m + n
 grid[i][j] 要么是 0 ，要么是 1 。


 Related Topics 广度优先搜索 图 数组 矩阵 最短路 堆（优先队列） 👍 10 👎 0

*/

/*
题型: 网格bfs
题目: 穿越网格图的安全路径
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findSafeWalk(grid [][]int, health int) bool {
	type pair struct{ x, y int }
	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}

	dis[0][0] = grid[0][0]
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		i, j := p.x, p.y
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				cost := grid[x][y]
				if dis[i][j]+cost < dis[x][y] {
					dis[x][y] = dis[i][j] + cost
					q[cost] = append(q[cost], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1] < health
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
