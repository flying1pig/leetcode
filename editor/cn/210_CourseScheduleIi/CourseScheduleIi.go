package main

/**
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。


 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。


 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。



 示例 1：


输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。


 示例 2：


输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

 示例 3：


输入：numCourses = 1, prerequisites = []
输出：[0]



提示：


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 ai != bi
 所有[ai, bi] 互不相同


 Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1020 👎 0

*/

/*
题型: 图论拓扑排序
题目: 课程表II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	//入度表，记录本课程的先休课程有几门
	inDegree := make([]int, numCourses)
	//出边表，记录本课程是哪些课程的先休
	outEdge := make([][]int, numCourses)
	// 结果集
	res := []int{}

	for i := 0; i < len(prerequisites); i++ {
		// 更新入度
		inDegree[prerequisites[i][0]]++
		// 更新出边
		outEdge[prerequisites[i][1]] = append(outEdge[prerequisites[i][1]], prerequisites[i][0])
	}

	// 队列用于记录入度为0的课程
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 没有不依赖其它课程的课程，返回空
	if len(queue) == 0 {
		return res
	}

	for len(queue) > 0 {
		// 取出队头，加入结果集
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)

		// 处理队头的出边, 把入度为0的入队
		for i := 0; i < len(outEdge[node]); i++ {
			out := outEdge[node][i]
			inDegree[out]--
			if inDegree[out] == 0 {
				queue = append(queue, out)
			}
		}
	}

	// 检查是否所有的课程都被安排，如果没有，说明有环，无法安排，返回空
	if len(res) == numCourses {
		return res
	}

	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
