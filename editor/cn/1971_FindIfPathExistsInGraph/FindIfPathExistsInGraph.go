package main

/**
有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中
edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。

 请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。

 给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回
true，否则返回 false 。



 示例 1：


输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
输出：true
解释：存在由顶点 0 到顶点 2 的路径:
- 0 → 1 → 2
- 0 → 2


 示例 2：


输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
输出：false
解释：不存在由顶点 0 到顶点 5 的路径.




 提示：


 1 <= n <= 2 * 10⁵
 0 <= edges.length <= 2 * 10⁵
 edges[i].length == 2
 0 <= ui, vi <= n - 1
 ui != vi
 0 <= source, destination <= n - 1
 不存在重复边
 不存在指向顶点自身的边


 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 229 👎 0

*/

/*
题型: 图论dfs
题目: 寻找图中是否存在路径
*/

// leetcode submit region begin(Prohibit modification and deletion)
func validPath(n int, edges [][]int, source int, destination int) bool {
	nodeMap := map[int][]int{}
	for _, edge := range edges {
		if nodes, ok := nodeMap[edge[0]]; ok {
			nodes = append(nodes, edge[1])
			nodeMap[edge[0]] = nodes
		} else {
			nodeMap[edge[0]] = []int{edge[1]}
		}

		if nodes, ok := nodeMap[edge[1]]; ok {
			nodes = append(nodes, edge[0])
			nodeMap[edge[1]] = nodes
		} else {
			nodeMap[edge[1]] = []int{edge[0]}
		}
	}

	var dfs func(i int) bool
	visited := map[int]interface{}{}
	dfs = func(i int) bool {
		if i == destination {
			return true
		}

		if nodes, ok := nodeMap[i]; ok {
			for _, node := range nodes {
				if _, ok := visited[node]; ok {
					continue
				}
				visited[node] = struct{}{}
				if dfs(node) {
					return true
				}
			}
		}
		return false
	}

	return dfs(source)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
