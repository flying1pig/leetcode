package main

/**
给你一个 m x n 的网格 grid。网格里的每个单元都代表一条街道。grid[i][j] 的街道可以是：


 1 表示连接左单元格和右单元格的街道。
 2 表示连接上单元格和下单元格的街道。
 3 表示连接左单元格和下单元格的街道。
 4 表示连接右单元格和下单元格的街道。
 5 表示连接左单元格和上单元格的街道。
 6 表示连接右单元格和上单元格的街道。




 你最开始从左上角的单元格 (0,0) 开始出发，网格中的「有效路径」是指从左上方的单元格 (0,0) 开始、一直到右下方的 (m-1,n-1) 结束的路径。该
路径必须只沿着街道走。

 注意：你 不能 变更街道。

 如果网格中存在有效的路径，则返回 true，否则返回 false 。



 示例 1：



 输入：grid = [[2,4,3],[6,5,2]]
输出：true
解释：如图所示，你可以从 (0, 0) 开始，访问网格中的所有单元格并到达 (m - 1, n - 1) 。


 示例 2：



 输入：grid = [[1,2,1],[1,2,1]]
输出：false
解释：如图所示，单元格 (0, 0) 上的街道没有与任何其他单元格上的街道相连，你只会停在 (0, 0) 处。


 示例 3：

 输入：grid = [[1,1,2]]
输出：false
解释：你会停在 (0, 1)，而且无法到达 (0, 2) 。


 示例 4：

 输入：grid = [[1,1,1,1,1,1,3]]
输出：true


 示例 5：

 输入：grid = [[2],[2],[2],[2],[2],[2],[6]]
输出：true




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 300
 1 <= grid[i][j] <= 6


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 91 👎 0

*/

/*
题型: 网格dfs
题目: 检查网格中是否存在有效路径
*/

// leetcode submit region begin(Prohibit modification and deletion)

func getRoute(x int) [][]int {
	if x == 1 {
		return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {0, -1, 1}, {0, -1, 4}, {0, -1, 6}}
	} else if x == 2 {
		return [][]int{{1, 0, 2}, {1, 0, 5}, {1, 0, 6}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
	} else if x == 3 {
		return [][]int{{0, -1, 1}, {0, -1, 4}, {0, -1, 6}, {1, 0, 2}, {1, 0, 5}, {1, 0, 6}}
	} else if x == 4 {
		return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {1, 0, 2}, {1, 0, 5}, {1, 0, 6}}
	} else if x == 5 {
		return [][]int{{0, -1, 1}, {0, -1, 4}, {0, -1, 6}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
	}
	return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
}

func hasValidPath(grid [][]int) bool {
	var (
		m       int
		n       int
		flag    bool
		iterate func(int, int)
	)
	m = len(grid)
	n = len(grid[0])
	if m == 1 && n == 1 {
		return true
	}
	iterate = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		route := getRoute(grid[x][y])
		temp := grid[x][y]
		grid[x][y] = 0
		for i := 0; i < len(route); i++ {
			newx := x + route[i][0]
			newy := y + route[i][1]
			if newx >= 0 && newx < m && newy >= 0 && newy < n && route[i][2] != grid[newx][newy] {
				continue
			}
			if newx == m-1 && newy == n-1 {
				flag = true
				return
			}
			iterate(newx, newy)
		}
		grid[x][y] = temp
	}
	flag = false
	iterate(0, 0)
	return flag
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
