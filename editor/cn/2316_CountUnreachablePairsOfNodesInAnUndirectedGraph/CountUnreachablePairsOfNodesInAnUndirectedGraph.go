package main

/**
给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [
ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。

 请你返回 无法互相到达 的不同 点对数目 。



 示例 1：



 输入：n = 3, edges = [[0,1],[0,2],[1,2]]
输出：0
解释：所有点都能互相到达，意味着没有点对无法互相到达，所以我们返回 0 。


 示例 2：



 输入：n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
输出：14
解释：总共有 14 个点对互相无法到达：
[[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6],[
5,6]]
所以我们返回 14 。




 提示：


 1 <= n <= 10⁵
 0 <= edges.length <= 2 * 10⁵
 edges[i].length == 2
 0 <= ai, bi < n
 ai != bi
 不会有重复边。


 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 136 👎 0

*/

/*
题型: 图论dfs
题目: 统计无向图中无法互相到达点对数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vis[x] = true // 避免重复访问同一个点
		size := 1
		for _, y := range g[x] {
			if !vis[y] {
				size += dfs(y)
			}
		}
		return size
	}

	total := 0
	for i, b := range vis {
		if !b { // 未访问的点：说明找到了一个新的连通块
			size := dfs(i)
			ans += int64(size) * int64(total)
			total += size
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
