package main

/**
你正在维护一个项目，该项目有 n 个方法，编号从 0 到 n - 1。

 给你两个整数 n 和 k，以及一个二维整数数组 invocations，其中 invocations[i] = [ai, bi] 表示方法 ai 调用了方法
bi。

 已知如果方法 k 存在一个已知的 bug。那么方法 k 以及它直接或间接调用的任何方法都被视为 可疑方法 ，我们需要从项目中移除这些方法。

 只有当一组方法没有被这组之外的任何方法调用时，这组方法才能被移除。

 返回一个数组，包含移除所有 可疑方法 后剩下的所有方法。你可以以任意顺序返回答案。如果无法移除 所有 可疑方法，则 不 移除任何方法。



 示例 1:


 输入: n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]


 输出: [0,1,2,3]

 解释:



 方法 2 和方法 1 是可疑方法，但它们分别直接被方法 3 和方法 0 调用。由于方法 3 和方法 0 不是可疑方法，我们无法移除任何方法，故返回所有方法。


 示例 2:


 输入: n = 5, k = 0, invocations = [[1,2],[0,2],[0,1],[3,4]]


 输出: [3,4]

 解释:



 方法 0、方法 1 和方法 2 是可疑方法，且没有被任何其他方法直接调用。我们可以移除它们。

 示例 3:


 输入: n = 3, k = 2, invocations = [[1,2],[0,1],[2,0]]


 输出: []

 解释:



 所有方法都是可疑方法。我们可以移除它们。



 提示:


 1 <= n <= 10⁵
 0 <= k <= n - 1
 0 <= invocations.length <= 2 * 10⁵
 invocations[i] == [ai, bi]
 0 <= ai, bi <= n - 1
 ai != bi
 invocations[i] != invocations[j]


 Related Topics 深度优先搜索 广度优先搜索 图 👍 8 👎 0

*/

/*
题型: 图论dfs
题目: 移除可疑的方法
*/

// leetcode submit region begin(Prohibit modification and deletion)
func remainingMethods(n, k int, invocations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, e := range invocations {
		g[e[0]] = append(g[e[0]], e[1])
	}

	// 标记所有可疑方法
	isSuspicious := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		isSuspicious[x] = true
		for _, y := range g[x] {
			if !isSuspicious[y] { // 避免无限递归
				dfs(y)
			}
		}
	}
	dfs(k)

	// 检查是否有【非可疑方法】->【可疑方法】的边
	for _, e := range invocations {
		if !isSuspicious[e[0]] && isSuspicious[e[1]] {
			// 无法移除可疑方法
			for i := range n {
				ans = append(ans, i)
			}
			return
		}
	}

	// 移除所有可疑方法
	for i, b := range isSuspicious {
		if !b {
			ans = append(ans, i)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
	从 k 开始 DFS 图，标记所有可能有 bug 的方法（节点）。题目把这些方法叫做可疑方法。
	遍历 invocations，如果存在从「非可疑方法」到「可疑方法」的边，则删除后无法编译，返回数组 [0,1,2,⋯,n−1]。
	否则可以正常删除，把非可疑方法加入答案。
*/

func main() {

}
