package main

/**
给你一棵有 n 个节点的无向树，节点编号为 0 到 n-1 ，它们中有一些节点有苹果。通过树上的一条边，需要花费 1 秒钟。你从 节点 0 出发，请你返回最少需
要多少秒，可以收集到所有苹果，并回到节点 0 。

 无向树的边由 edges 给出，其中 edges[i] = [fromi, toi] ，表示有一条边连接 from 和 toi 。除此以外，还有一个布尔数组
hasApple ，其中 hasApple[i] = true 代表节点 i 有一个苹果，否则，节点 i 没有苹果。



 示例 1：




输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
false,true,false,true,true,false]
输出：8
解释：上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。


 示例 2：




输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
false,true,false,false,true,false]
输出：6
解释：上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。


 示例 3：


输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
false,false,false,false,false,false]
输出：0




 提示：


 1 <= n <= 10^5
 edges.length == n - 1
 edges[i].length == 2
 0 <= ai < bi <= n - 1
 hasApple.length == n


 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 👍 108 👎 0

*/

/*
题型: 一般树
题目: 收集树上所有苹果的最少时间
*/

// leetcode submit region begin(Prohibit modification and deletion)

func minTime(n int, edges [][]int, hasApple []bool) int {
	set, f := make([][]int, n), false // set记录的是每个点对应的边，f记录这棵树有没有苹果
	for _, e := range edges {
		f = f || hasApple[e[0]] || hasApple[e[1]]
		set[e[0]] = append(set[e[0]], e[1])
		set[e[1]] = append(set[e[1]], e[0])
	}
	if !f {
		return 0
	}
	fs := make([]bool, n) //fs记录已经遍历过的点
	var dfs func(i int) int
	dfs = func(i int) int { // 计算从i点出发回到i点的有效步数
		r := 0
		fs[i] = true // 把当前点置为true
		for _, c := range set[i] {
			if fs[c] {
				continue
			}
			if hasApple[c] {
				r += 2
			} // 如果当前点有苹果直接+2（回溯是有效的）
			x := dfs(c)
			if !hasApple[c] && x != 0 { // 如果当前点没有苹果但是它的子树有苹果同样+2（回溯是有效的）
				r += 2
			}
			r += x
		}
		return r
	}
	return dfs(0)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
