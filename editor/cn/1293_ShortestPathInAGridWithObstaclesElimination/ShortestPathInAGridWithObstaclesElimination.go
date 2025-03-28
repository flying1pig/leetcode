package main

/**
给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。

 如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路
径，则返回 -1 。



 示例 1：




输入： grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
输出：6
解释：
不消除任何障碍的最短路径是 10。
消除位置 (3,2) 处的障碍后，最短路径是 6 。该路径是 (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2
) -> (4,2).


 示例 2：




输入：grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
输出：-1
解释：我们至少需要消除两个障碍才能找到这样的路径。




 提示：


 grid.length == m
 grid[0].length == n
 1 <= m, n <= 40
 1 <= k <= m*n
 grid[i][j] 是 0 或 1
 grid[0][0] == grid[m-1][n-1] == 0


 Related Topics 广度优先搜索 数组 矩阵 👍 291 👎 0

*/

/*
题型: 网格bfs
题目: 网格中的最短路径
*/

// leetcode submit region begin(Prohibit modification and deletion)
func shortestPath(grid [][]int, k int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	// 四个方向
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 定义vis数组，记录访问过的数字，初始化置-1代表未访问
	// vis数组在BFS时还有另外一层含义，即当前该位置剩余步数，初始化[0， 0]位置即为k
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vis[i][j] = -1
		}
	}
	vis[0][0] = k
	queue := [][]int{}
	queue = append(queue, []int{0, 0})
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x := temp[0] + dir[0]
				y := temp[1] + dir[1]
				// 判断当前位置的合法性，如果当前位置未访问，或者相比之前其他位置访问该位置时剩余步数更多，则进入
				if x >= 0 && x < m && y >= 0 && y < n && (vis[x][y] == -1 || vis[temp[0]][temp[1]]-grid[x][y] > vis[x][y]) {
					// 判断当前位置是否还有足够的步数进入，如果当前位置为0，则一定可以进入，如果当前位置为，则至少还需要一步才能进入
					if grid[x][y] == 0 || grid[x][y] == 1 && vis[temp[0]][temp[1]] > 0 {
						if x == m-1 && y == n-1 {
							return res + 1
						}
						queue = append(queue, []int{x, y})
						// 走到当前位置后，当前位置剩余的步数判断
						if grid[x][y] == 1 {
							vis[x][y] = vis[temp[0]][temp[1]] - 1
						} else {
							vis[x][y] = vis[temp[0]][temp[1]]
						}
					}
				}
			}
		}
		res++
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
