package main

/**
这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k 。

 给定三个整数 n、k 和 target，请返回投掷骰子的所有可能得到的结果（共有 kⁿ 种方式），使得骰子面朝上的数字总和等于 target。

 由于答案可能很大，你需要对 10⁹ + 7 取模。



 示例 1：


输入：n = 1, k = 6, target = 3
输出：1
解释：你掷了一个有 6 个面的骰子。
得到总和为 3 的结果的方式只有一种。


 示例 2：


输入：n = 2, k = 6, target = 7
输出：6
解释：你掷了两个骰子，每个骰子有 6 个面。
有 6 种方式得到总和为 7 的结果: 1+6, 2+5, 3+4, 4+3, 5+2, 6+1。


 示例 3：


输入：n = 30, k = 30, target = 500
输出：222616187
解释：返回的结果必须对 10⁹ + 7 取模。



 提示：


 1 <= n, k <= 30
 1 <= target <= 1000


 Related Topics 动态规划 👍 303 👎 0

*/

/*
题型: dp
题目: 掷骰子等于目标和的方法数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numRollsToTarget(n, k, target int) int {
	if target < n || target > n*k {
		return 0 // 无法组成 target
	}
	const mod = 1_000_000_007
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, target-n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if j == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		for x := 0; x < k && x <= j; x++ { // 掷出了 x
			res = (res + dfs(i-1, j-x)) % mod
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n, target-n)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
dfs(i,j)=dfs(i−1,j)+dfs(i−1,j−1)+⋯+dfs(i−1,j−min(k−1,j))
*/

/*
记忆化搜索
func numRollsToTarget(n, k, target int) int {
    if target < n || target > n*k {
        return 0 // 无法组成 target
    }
    const mod = 1_000_000_007
    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, target-n+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i == 0 {
            if j == 0 {
                return 1
            }
            return 0
        }
        p := &memo[i][j]
        if *p != -1 { // 之前计算过
            return *p
        }
        res := 0
        for x := 0; x < k && x <= j; x++ { // 掷出了 x
            res = (res + dfs(i-1, j-x)) % mod
        }
        *p = res // 记忆化
        return res
    }
    return dfs(n, target-n)
}

*/

/*
	func numRollsToTarget(n, k, target int) int {
	    if target < n || target > n*k {
	        return 0 // 无法组成 target
	    }
	    const mod = 1_000_000_007
	    f := make([][]int, n+1)
	    for i := range f {
	        f[i] = make([]int, target-n+1)
	    }
	    f[0][0] = 1
	    for i := 1; i <= n; i++ {
	        for j := 0; j <= target-n; j++ {
	            for x := 0; x < k && x <= j; x++ { // 掷出了 x
	                f[i][j] = (f[i][j] + f[i-1][j-x]) % mod
	            }
	        }
	    }
	    return f[n][target-n]
	}
*/
func main() {

}
