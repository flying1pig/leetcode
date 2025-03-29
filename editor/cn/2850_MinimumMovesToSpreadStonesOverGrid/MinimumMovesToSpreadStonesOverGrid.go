package main

import "math"

/**
给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有 多
个 石头。

 每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。

 请你返回每个格子恰好有一个石头的 最少移动次数 。



 示例 1：




输入：grid = [[1,1,0],[1,1,1],[1,2,1]]
输出：3
解释：让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
总共需要 3 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 3 。


 示例 2：




输入：grid = [[1,3,0],[1,0,0],[1,0,3]]
输出：4
解释：让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (0,1) 移动到 (0,2) 。
2 - 将一个石头从格子 (0,1) 移动到 (1,1) 。
3 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
4 - 将一个石头从格子 (2,2) 移动到 (2,1) 。
总共需要 4 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 4 。




 提示：


 grid.length == grid[i].length == 3
 0 <= grid[i][j] <= 9
 grid 中元素之和为 9 。


 Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 76 👎 0

*/

/*
题型: 回溯
题目: 将石头分散到网格图的最少移动次数
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	var from, to []pair
	for i, row := range grid {
		for j, cnt := range row {
			if cnt > 1 {
				for k := 1; k < cnt; k++ {
					from = append(from, pair{i, j})
				}
			} else if cnt == 0 {
				to = append(to, pair{i, j})
			}
		}
	}

	ans := math.MaxInt
	permute(from, 0, func() {
		total := 0
		for i, f := range from {
			total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
		}
		ans = min(ans, total)
	})
	return ans
}

func permute(a []pair, i int, do func()) {
	if i == len(a) {
		do()
		return
	}
	permute(a, i+1, do)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, i+1, do)
		a[i], a[j] = a[j], a[i]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
