package main

/**
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如
果要学习课程 ai 则 必须 先学习课程 bi 。


 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。


 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



 示例 1：


输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

 示例 2：


输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。



 提示：


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= 5000
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 prerequisites[i] 中的所有课程对 互不相同


 Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 2118 👎 0

*/

/*
题型: 图论dfs，三色标记法判环
题目: 课程表
*/

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 // x 正在访问中
		for _, y := range g[x] {
			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
				return true // 找到了环
			}
		}
		colors[x] = 2 // x 完全访问完毕
		return false  // 没有找到环
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false // 有环
		}
	}
	return true // 没有环
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
	1. 建图：把每个 prerequisites[i]=[a,b] 看成一条有向边 b→a，构建一个有向图 g。
	2. 创建长为 numCourses 的颜色数组 colors，所有元素值初始化成 0。
	3. 遍历 colors，如果 colors[i]=0，则调用递归函数 dfs(i)。
	4. 执行 dfs(x)：
		a. 首先标记 colors[x]=1，表示节点 x 正在访问中。
		b. 然后遍历 x 的邻居 y。如果 colors[y]=1，则找到环，返回 true。
如果 colors[y]=0（没有访问过）且 dfs(y) 返回了 true，那么 dfs(x) 也返回 true。
		c. 如果没有找到环，那么先标记 colors[x]=2，表示 x 已经完全访问完毕，然后返回 false。
	5. 如果 dfs(i) 返回 true，那么找到了环，返回 false。
	6. 如果遍历完所有节点也没有找到环，返回 true

*/

func main() {

}
