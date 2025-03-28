package main

import (
	"math"
	"slices"
)

/**
有 n 个网络节点，标记为 1 到 n。

 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi
是一个信号从源节点传递到目标节点的时间。

 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。



 示例 1：




输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2


 示例 2：


输入：times = [[1,2,1]], n = 2, k = 1
输出：1


 示例 3：


输入：times = [[1,2,1]], n = 2, k = 2
输出：-1




 提示：


 1 <= k <= n <= 100
 1 <= times.length <= 6000
 times[i].length == 3
 1 <= ui, vi <= n
 ui != vi
 0 <= wi <= 100
 所有 (ui, vi) 对都 互不相同（即，不含重复边）


 Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 840 👎 0

*/

/*
题型: Dijkstra
题目: 网络延迟时间
*/

// leetcode submit region begin(Prohibit modification and deletion)
func networkDelayTime(times [][]int, n, k int) int {
	const inf = math.MaxInt / 2 // 防止加法溢出
	g := make([][]int, n)       // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		g[t[0]-1][t[1]-1] = t[2]
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 {
			return slices.Max(dis)
		}
		if dis[x] == inf { // 有节点无法到达
			return -1
		}
		done[x] = true // 最短路长度已确定（无法变得更小）
		for y, d := range g[x] {
			// 更新 x 的邻居的最短路
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/*
写法一: 朴素dijkstra(适用于稠密图)
func networkDelayTime(times [][]int, n, k int) int {
	const inf = math.MaxInt / 2 // 防止加法溢出
	g := make([][]int, n)       // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		g[t[0]-1][t[1]-1] = t[2]
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 {
			return slices.Max(dis)
		}
		if dis[x] == inf { // 有节点无法到达
			return -1
		}
		done[x] = true // 最短路长度已确定（无法变得更小）
		for y, d := range g[x] {
			// 更新 x 的邻居的最短路
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
}
*/

/*
堆优化dijkstra(适用于稀疏图)
func networkDelayTime(times [][]int, n, k int) int {
    type edge struct{ to, wt int }
    g := make([][]edge, n) // 邻接表
    for _, t := range times {
        g[t[0]-1] = append(g[t[0]-1], edge{t[1] - 1, t[2]})
    }

    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt
    }
    dis[k-1] = 0
    h := hp{{0, k - 1}}
    for len(h) > 0 {
        p := heap.Pop(&h).(pair)
        dx := p.dis
        x := p.x
        if dx > dis[x] { // x 之前出堆过
            continue
        }
        for _, e := range g[x] {
            y := e.to
            newDis := dx + e.wt
            if newDis < dis[y] {
                dis[y] = newDis // 更新 x 的邻居的最短路
                heap.Push(&h, pair{newDis, y})
            }
        }
    }
    mx := slices.Max(dis)
    if mx < math.MaxInt {
        return mx
    }
    return -1
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
*/

func main() {

}

/*
其余dijkstra算法题单:
3341. 到达最后一个房间的最少时间 I 1721 网格图
3112. 访问消失节点的最少时间 1757 理解原理
2642. 设计可以求最短路径的图类 1811
1514. 概率最大的路径 1846
3342. 到达最后一个房间的最少时间 II 1862 网格图
1631. 最小体力消耗路径 1948 网格图 做法不止一种
1786. 从第一个节点出发到最后一个节点的受限路径数 2079
3123. 最短路径中的边 2093
1976. 到达目的地的方案数 2095
778. 水位上升的泳池中游泳 2097 网格图 做法不止一种
2662. 前往目标的最小代价 2154
3377. 使两个整数相等的数位操作 2186
*/
