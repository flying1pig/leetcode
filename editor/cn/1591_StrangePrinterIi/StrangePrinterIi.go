package main

/**
给你一个奇怪的打印机，它有如下两个特殊的打印规则：


 每一次操作时，打印机会用同一种颜色打印一个矩形的形状，每次打印会覆盖矩形对应格子里原本的颜色。
 一旦矩形根据上面的规则使用了一种颜色，那么 相同的颜色不能再被使用 。


 给你一个初始没有颜色的 m x n 的矩形 targetGrid ，其中 targetGrid[row][col] 是位置 (row, col) 的颜色。

 如果你能按照上述规则打印出矩形 targetGrid ，请你返回 true ，否则返回 false 。



 示例 1：



 输入：targetGrid = [[1,1,1,1],[1,2,2,1],[1,2,2,1],[1,1,1,1]]
输出：true


 示例 2：



 输入：targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,4],[5,5,1,4]]
输出：true


 示例 3：

 输入：targetGrid = [[1,2,1],[2,1,2],[1,2,1]]
输出：false
解释：没有办法得到 targetGrid ，因为每一轮操作使用的颜色互不相同。

 示例 4：

 输入：targetGrid = [[1,1,1],[3,1,3]]
输出：false




 提示：


 m == targetGrid.length
 n == targetGrid[i].length
 1 <= m, n <= 60
 1 <= targetGrid[row][col] <= 60


 Related Topics 图 拓扑排序 数组 矩阵 👍 62 👎 0

*/

/*
题型: 图论拓扑排序
题目: 奇怪的打印机 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func isPrintable(targetGrid [][]int) bool {

	// colors: 每个颜色的左上角坐标和右下角坐标
	colors := make(map[int][]int)
	for i := 0; i < len(targetGrid); i++ {
		for j := 0; j < len(targetGrid[i]); j++ {
			if pos, ok := colors[targetGrid[i][j]]; ok {
				pos[0] = min(i, pos[0])
				pos[1] = min(j, pos[1])
				pos[2] = max(i, pos[2])
				pos[3] = max(j, pos[3])
			} else {
				colors[targetGrid[i][j]] = []int{i, j, i, j}
			}
		}
	}

	// 检查颜色color是否构成矩形
	check := func(color int) bool {
		pos := colors[color]
		for i := pos[0]; i <= pos[2]; i++ {
			for j := pos[1]; j <= pos[3]; j++ {
				if targetGrid[i][j] != color && targetGrid[i][j] != 0 {
					return false
				}
			}
		}
		return true
	}

	// 把颜色为color的块置0
	mark := func(color int) {
		for i := 0; i < len(targetGrid); i++ {
			for j := 0; j < len(targetGrid[i]); j++ {
				if targetGrid[i][j] == color {
					targetGrid[i][j] = 0
				}
			}
		}
	}

	// 对每种颜色，检查是否是矩形，是则把相应块置0并删除该颜色
	keepGoing := true
	for keepGoing {
		keepGoing = false
		for color, _ := range colors {
			if check(color) {
				mark(color)
				delete(colors, color)
				keepGoing = true
				break
			}
		}
	}

	// 检查是否剩余颜色
	return len(colors) == 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
