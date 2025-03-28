package main

/**
你现在手里有一份大小为
 n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地。

 请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 -1。

 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 - x1
| + |y0 - y1| 。



 示例 1：




输入：grid = [[1,0,1],[0,0,0],[1,0,1]]
输出：2
解释：
海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。


 示例 2：




输入：grid = [[1,0,0],[0,0,0],[0,0,0]]
输出：4
解释：
海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。




 提示：





 n == grid.length
 n == grid[i].length
 1 <= n <= 100
 grid[i][j] 不是 0 就是 1


 Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 410 👎 0

*/

/*
题型: 网格bfs
题目: 地图分析
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(grid [][]int) int {
	var dx = []int{0, 1, -1, 0}
	var dy = []int{1, 0, 0, -1}
	q := [][]int{}
	dis := make([][]int, len(grid))
	for i := range dis {
		dis[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	ans := -1
	for len(q) > 0 {
		for sz := len(q); sz > 0; sz-- { // 多源 BFS 的搜索模板
			t := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				x, y := dx[i]+t[0], dy[i]+t[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 1 {
					continue
				}
				if dis[x][y] != 0 { // 如果已经有陆地走到这里, 更新一下最大值
					ans = max(ans, dis[x][y])
					continue
				}
				dis[x][y] = dis[t[0]][t[1]] + 1
				q = append(q, []int{x, y})
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
	陆地出发搜海洋
	dis 数组存放的就是曼哈顿距离
	使用多源 BFS 搜索，每个陆地同时走
	如果陆地同时走到一块地上，更新距离的最大值并返回即可
*/

func main() {

}
