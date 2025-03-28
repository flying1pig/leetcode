package main

/**
给定一个整数 n，即有向图中的节点数，其中节点标记为 0 到 n - 1。图中的每条边为红色或者蓝色，并且可能存在自环或平行边。

 给定两个数组 redEdges 和 blueEdges，其中：


 redEdges[i] = [ai, bi] 表示图中存在一条从节点 ai 到节点 bi 的红色有向边，
 blueEdges[j] = [uj, vj] 表示图中存在一条从节点 uj 到节点 vj 的蓝色有向边。


 返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么
 answer[x] = -1。



 示例 1：


输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
输出：[0,1,-1]


 示例 2：


输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
输出：[0,1,-1]




 提示：


 1 <= n <= 100
 0 <= redEdges.length, blueEdges.length <= 400
 redEdges[i].length == blueEdges[j].length == 2
 0 <= ai, bi, uj, vj < n


 Related Topics 广度优先搜索 图 👍 326 👎 0

*/

/*
题型: 图论bfs
题目: 颜色交替的最短路径
*/

// leetcode submit region begin(Prohibit modification and deletion)
func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	g := make([][]pair, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if dis[x] < 0 {
				dis[x] = level
			}
			for _, to := range g[x] {
				if to.color != p.color && !vis[to.x][to.color] {
					vis[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return dis
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
