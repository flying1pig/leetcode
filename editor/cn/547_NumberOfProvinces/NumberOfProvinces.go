package main

/**


 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。




 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
isConnected[i][j] = 0 表示二者不直接相连。

 返回矩阵中 省份 的数量。



 示例 1：


输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2


 示例 2：


输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3




 提示：


 1 <= n <= 200
 n == isConnected.length
 n == isConnected[i].length
 isConnected[i][j] 为 1 或 0
 isConnected[i][i] == 1
 isConnected[i][j] == isConnected[j][i]


 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 1194 👎 0

*/

/*
题型: 图论dfs
题目: 省份数量
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
深度优先搜索
func findCircleNum(isConnected [][]int) (ans int) {
    vis := make([]bool, len(isConnected))
    var dfs func(int)
    dfs = func(from int) {
        vis[from] = true
        for to, conn := range isConnected[from] {
            if conn == 1 && !vis[to] {
                dfs(to)
            }
        }
    }
    for i, v := range vis {
        if !v {
            ans++
            dfs(i)
        }
    }
    return
}
*/

/*
广度优先搜索
func findCircleNum(isConnected [][]int) (ans int) {
    vis := make([]bool, len(isConnected))
    for i, v := range vis {
        if !v {
            ans++
            queue := []int{i}
            for len(queue) > 0 {
                from := queue[0]
                queue = queue[1:]
                vis[from] = true
                for to, conn := range isConnected[from] {
                    if conn == 1 && !vis[to] {
                        queue = append(queue, to)
                    }
                }
            }
        }
    }
    return
}
*/

/*
并查集

	func findCircleNum(isConnected [][]int) (ans int) {
	    n := len(isConnected)
	    parent := make([]int, n)
	    for i := range parent {
	        parent[i] = i
	    }
	    var find func(int) int
	    find = func(x int) int {
	        if parent[x] != x {
	            parent[x] = find(parent[x])
	        }
	        return parent[x]
	    }
	    union := func(from, to int) {
	        parent[find(from)] = find(to)
	    }

	    for i, row := range isConnected {
	        for j := i + 1; j < n; j++ {
	            if row[j] == 1 {
	                union(i, j)
	            }
	        }
	    }
	    for i, p := range parent {
	        if i == p {
	            ans++
	        }
	    }
	    return
	}
*/
func main() {

}
