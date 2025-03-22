package main

/**
给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][n -
 1]）。机器人每次只能向下或者向右移动一步。

 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。

 返回机器人能够到达右下角的不同路径数量。

 测试用例保证答案小于等于 2 * 10⁹。



 示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右


 示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1




 提示：


 m == obstacleGrid.length
 n == obstacleGrid[i].length
 1 <= m, n <= 100
 obstacleGrid[i][j] 为 0 或 1


 👍 1397 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	mem := make([]int, n+1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 && obstacleGrid[0][0] == 0 {
				mem[1] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				mem[j+1] = 0
				continue
			}
			mem[j+1] = mem[j] + mem[j+1]
		}
	}
	return mem[n]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义f(i,j)为从(0.0)到(i,j)的路径和
如果(i,j) = 1, f(i,j) = 0
从左边来，f(i,j) = f(i,j-1)，条件是(i,j-1) == 0
从上边来，f(i,j) = f(i-1,j)，条件是(i-1,j) == 0
*/

/*
递推:
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m,n := len(obstacleGrid),len(obstacleGrid[0])
	mem := make([][]int,m+1)
	for i:=0;i<=m;i++ {
		mem[i] = make([]int,n+1)
	}
	for i:=0;i<m;i++ {
		for j :=0;j<n;j++ {
			if i == 0 && j == 0 && obstacleGrid[0][0] == 0{
				mem[1][1] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				continue
			}
			mem[i+1][j+1] = mem[i+1][j] + mem[i][j+1]
		}
	}
	return mem[m][n]
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
空间优化:
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m,n := len(obstacleGrid),len(obstacleGrid[0])
	mem := make([]int,n+1)

	for i:=0;i<m;i++ {
		for j :=0;j<n;j++ {
			if i == 0 && j == 0 && obstacleGrid[0][0] == 0{
				mem[1] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				mem[j+1] = 0
				continue
			}
			mem[j+1] = mem[j] + mem[j+1]
		}
	}
	return mem[n]
}
时间复杂度: o(mn)
空间复杂度: o(n)
*/

func main() {

}
