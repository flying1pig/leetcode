package main

/**
给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

 两个相邻元素间的距离为 1 。



 示例 1：




输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0,0],[0,1,0],[0,0,0]]


 示例 2：




输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
输出：[[0,0,0],[0,1,0],[1,2,1]]




 提示：


 m == mat.length
 n == mat[i].length
 1 <= m, n <= 10⁴
 1 <= m * n <= 10⁴
 mat[i][j] is either 0 or 1.
 mat 中至少有一个 0


 Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 986 👎 0

*/

/*
题型: 网格bfs
题目: 01 矩阵
*/

// leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[0]))
	}

	zero := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				zero = append(zero, []int{i, j})
			}
		}
	}
	moves := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for len(zero) != 0 {
		l := len(zero)
		for i := 0; i < l; i++ {
			r1 := zero[0][0]
			c1 := zero[0][1]
			zero = zero[1:len(zero)]
			for j := 0; j < len(moves); j++ {
				r2 := r1 + moves[j][0]
				c2 := c1 + moves[j][1]
				if inArea(mat, r2, c2) && mat[r2][c2] == 1 {
					zero = append(zero, []int{r2, c2})
					mat[r2][c2] = 2
					res[r2][c2] = res[r1][c1] + 1
				}
			}
		}
	}
	return res
}

func inArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

//leetcode submit region end(Prohibit modification and deletion)

/*
多源BFS,每一层所到的点是两点之间的最短距离
需要一个queue记录每一层的节点
需要额外同样大小的矩阵记录距离
*/
func main() {

}
