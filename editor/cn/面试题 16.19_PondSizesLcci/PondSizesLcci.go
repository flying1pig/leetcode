package main

import "sort"

/**
你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连
接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
 示例：
 输入：
[
  [0,2,1,0],
  [0,1,0,1],
  [1,1,0,1],
  [0,1,0,1]
]
输出： [1,2,4]

 提示：

 0 < len(land) <= 1000
 0 < len(land[i]) <= 1000


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 163 👎 0

*/

/*
题型: 网格dfs
题目: 水域大小
*/

// leetcode submit region begin(Prohibit modification and deletion)
func pondSizes(land [][]int) (ans []int) {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || land[x][y] != 0 {
			return 0
		}
		land[x][y] = 1 // 标记 (x,y) 被访问，避免重复访问
		cnt0 := 1
		for i := x - 1; i <= x+1; i++ { // 访问八方向的 0
			for j := y - 1; j <= y+1; j++ {
				cnt0 += dfs(i, j)
			}
		}
		return cnt0
	}
	for i, row := range land {
		for j, x := range row {
			if x == 0 { // 从没有访问过的 0 出发
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	return
}

//时间复杂度：O(mn)
//空间复杂度：O(mn)
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
