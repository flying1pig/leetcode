package main

/**
欢迎各位勇者来到力扣城，本次试炼主题为「信物传送」。

 本次试炼场地设有若干传送带，matrix[i][j] 表示第 i 行 j 列的传送带运作方向，"^","v","<",">" 这四种符号分别表示 上、下、左、
右 四个方向。信物会随传送带的方向移动。勇者每一次施法操作，可临时变更一处传送带的方向，在物品经过后传送带恢复原方向。

 通关信物初始位于坐标 start处，勇者需要将其移动到坐标 end 处，请返回勇者施法操作的最少次数。

 注意：


 start 和 end 的格式均为 [i,j]


 示例 1：


 输入：matrix = [">>v","v^<","<><"], start = [0,1], end = [2,0]


 输出：1

 解释： 如上图所示 当信物移动到 [1,1] 时，勇者施法一次将 [1,1] 的传送方向 ^ 从变更为 < 从而信物移动到 [1,0]，后续到达 end 位置
 因此勇者最少需要施法操作 1 次

 示例 2：


 输入：matrix = [">>v",">>v","^<<"], start = [0,0], end = [1,1]


 输出：0

 解释：勇者无需施法，信物将自动传送至 end 位置

 示例 3：


 输入：matrix = [">^^>","<^v>","^v^<"], start = [0,0], end = [1,3]


 输出：3

 提示：


 matrix 中仅包含 '^'、'v'、'<'、'>'
 0 < matrix.length <= 100
 0 < matrix[i].length <= 100
 0 <= start[0],end[0] < matrix.length
 0 <= start[1],end[1] < matrix[i].length


 Related Topics 广度优先搜索 图 数组 矩阵 最短路 堆（优先队列） 👍 28 👎 0

*/

/*
题型: 网格bfs
题目: 信物传送
*/

// leetcode submit region begin(Prohibit modification and deletion)
func conveyorBelt(matrix []string, start []int, end []int) int {
	n := len(matrix)
	m := len(matrix[0])
	total := n * m
	dst := make([][]int, n)
	for i := range dst {
		dst[i] = make([]int, m)
		for j := range dst[i] {
			dst[i][j] = total
		}
	}
	type point struct{ x, y int }
	var dirs4 = []point{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1}}
	arrowCorrect := "v<^>"
	queue := []*point{{end[0], end[1]}}
	dst[end[0]][end[1]] = 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		ux := p.x
		uy := p.y
		for i := 0; i < 4; i++ {
			nx := ux + dirs4[i].x
			ny := uy + dirs4[i].y
			if nx >= 0 && ny >= 0 && nx < n && ny < m {
				if matrix[nx][ny] == arrowCorrect[i] {
					if dst[nx][ny] > dst[ux][uy] {
						dst[nx][ny] = dst[ux][uy]
						queue = append(queue, &point{nx, ny})
					}
				} else {
					if dst[nx][ny] > dst[ux][uy]+1 {
						dst[nx][ny] = dst[ux][uy] + 1
						queue = append(queue, &point{nx, ny})
					}
				}
			}
		}
	}
	return dst[start[0]][start[1]]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
