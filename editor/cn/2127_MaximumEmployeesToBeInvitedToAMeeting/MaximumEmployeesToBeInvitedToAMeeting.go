package main

/**
一个公司准备组织一场会议，邀请名单上有 n 位员工。公司准备了一张 圆形 的桌子，可以坐下 任意数目 的员工。

 员工编号为 0 到 n - 1 。每位员工都有一位 喜欢 的员工，每位员工 当且仅当 他被安排在喜欢员工的旁边，他才会参加会议。每位员工喜欢的员工 不会 是他
自己。

 给你一个下标从 0 开始的整数数组 favorite ，其中 favorite[i] 表示第 i 位员工喜欢的员工。请你返回参加会议的 最多员工数目 。



 示例 1：




输入：favorite = [2,2,1,2]
输出：3
解释：
上图展示了公司邀请员工 0，1 和 2 参加会议以及他们在圆桌上的座位。
没办法邀请所有员工参与会议，因为员工 2 没办法同时坐在 0，1 和 3 员工的旁边。
注意，公司也可以邀请员工 1，2 和 3 参加会议。
所以最多参加会议的员工数目为 3 。


 示例 2：


输入：favorite = [1,2,0]
输出：3
解释：
每个员工都至少是另一个员工喜欢的员工。所以公司邀请他们所有人参加会议的前提是所有人都参加了会议。
座位安排同图 1 所示：
- 员工 0 坐在员工 2 和 1 之间。
- 员工 1 坐在员工 0 和 2 之间。
- 员工 2 坐在员工 1 和 0 之间。
参与会议的最多员工数目为 3 。


 示例 3：




输入：favorite = [3,0,1,4,1]
输出：4
解释：
上图展示了公司可以邀请员工 0，1，3 和 4 参加会议以及他们在圆桌上的座位。
员工 2 无法参加，因为他喜欢的员工 1 旁边的座位已经被占领了。
所以公司只能不邀请员工 2 。
参加会议的最多员工数目为 4 。




 提示：


 n == favorite.length
 2 <= n <= 10⁵
 0 <= favorite[i] <= n - 1
 favorite[i] != i


 Related Topics 深度优先搜索 图 拓扑排序 👍 230 👎 0

*/

/*
题型: 基环树
题目: 参加会议的最多员工数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	deg := make([]int, n)
	for _, f := range favorite {
		deg[f]++ // 统计基环树每个节点的入度
	}

	rg := make([][]int, n) // 反图
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 拓扑排序，剪掉图上所有树枝
		x := q[0]
		q = q[1:]
		y := favorite[x] // x 只有一条出边
		rg[y] = append(rg[y], x)
		if deg[y]--; deg[y] == 0 {
			q = append(q, y)
		}
	}

	// 通过反图 rg 寻找树枝上最深的链
	var rdfs func(int) int
	rdfs = func(x int) int {
		maxDepth := 1
		for _, son := range rg[x] {
			maxDepth = max(maxDepth, rdfs(son)+1)
		}
		return maxDepth
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range deg {
		if d == 0 {
			continue
		}

		// 遍历基环上的点
		deg[i] = 0    // 将基环上的点的入度标记为 0，避免重复访问
		ringSize := 1 // 基环长度
		for x := favorite[i]; x != i; x = favorite[x] {
			deg[x] = 0 // 将基环上的点的入度标记为 0，避免重复访问
			ringSize++
		}

		if ringSize == 2 { // 基环长度为 2
			sumChainSize += rdfs(i) + rdfs(favorite[i]) // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环长度的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}

/*
基环树是一个连通图，有n个点和n条边，相比于树的n个点与n−1 条边，会多出一个环（有且只有一个环），
所以称为基环树。
多颗基环树也符合n 个点和n 条边的定义，由多棵基环树组成的森林称作基环森林。
有向的基环树还可以分为 in-tree（内向树） 和 out-tree（外向树），分别是每个点出度为1 以及
每个点入度为1的基环树，直观的可以看出内向树的边都指向里，外向树的边都指向外。
*/
