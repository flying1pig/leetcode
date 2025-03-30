package main

import "strings"

/**
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。


 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。


 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。



 示例 1：


输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n
的值 3 。


 示例 2：


输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。




 提示：


 1 <= strs.length <= 600
 1 <= strs[i].length <= 100
 strs[i] 仅由 '0' 和 '1' 组成
 1 <= m, n <= 100


 Related Topics 数组 字符串 动态规划 👍 1231 👎 0

*/

/*
题型: dp
题目: 一和零
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxForm(strs []string, m, n int) int {
	k := len(strs)
	cnt0 := make([]int, k)
	for i, s := range strs {
		cnt0[i] = strings.Count(s, "0")
	}

	memo := make([][][]int, k)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := dfs(i-1, j, k) // 不选 strs[i]
		cnt1 := len(strs[i]) - cnt0[i]
		if j >= cnt0[i] && k >= cnt1 {
			res = max(res, dfs(i-1, j-cnt0[i], k-cnt1)+1) // 选 strs[i]
		}
		*p = res // 记忆化
		return res
	}
	return dfs(k-1, m, n)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程, 考虑strs[i]选或不选
	不选:dfs(i,j,k)=dfs(i−1,j,k)
	选: dfs(i,j,k)=dfs(i−1,j−cnt0[i],k−cnt1[i])+1
边界条件:
	dfs(-1,j,k) = 0
*/

/*
记忆化搜索

	func findMaxForm(strs []string, m, n int) int {
	    k := len(strs)
	    cnt0 := make([]int, k)
	    for i, s := range strs {
	        cnt0[i] = strings.Count(s, "0")
	    }

	    memo := make([][][]int, k)
	    for i := range memo {
	        memo[i] = make([][]int, m+1)
	        for j := range memo[i] {
	            memo[i][j] = make([]int, n+1)
	            for k := range memo[i][j] {
	                memo[i][j][k] = -1 // -1 表示没有计算过
	            }
	        }
	    }
	    var dfs func(int, int, int) int
	    dfs = func(i, j, k int) int {
	        if i < 0 {
	            return 0
	        }
	        p := &memo[i][j][k]
	        if *p != -1 { // 之前计算过
	            return *p
	        }
	        res := dfs(i-1, j, k) // 不选 strs[i]
	        cnt1 := len(strs[i]) - cnt0[i]
	        if j >= cnt0[i] && k >= cnt1 {
	            res = max(res, dfs(i-1, j-cnt0[i], k-cnt1)+1) // 选 strs[i]
	        }
	        *p = res // 记忆化
	        return res
	    }
	    return dfs(k-1, m, n)
	}
*/

/*
递推
func findMaxForm(strs []string, m, n int) int {
    k := len(strs)
    f := make([][][]int, k+1)
    for i := range f {
        f[i] = make([][]int, m+1)
        for j := range f[i] {
            f[i][j] = make([]int, n+1)
        }
    }
    for i, s := range strs {
        cnt0 := strings.Count(s, "0")
        cnt1 := len(s) - cnt0
        for j := range m + 1 {
            for k := range n + 1 {
                if j >= cnt0 && k >= cnt1 {
                    f[i+1][j][k] = max(f[i][j][k], f[i][j-cnt0][k-cnt1]+1)
                } else {
                    f[i+1][j][k] = f[i][j][k]
                }
            }
        }
    }
    return f[k][m][n]
}
*/

/*
空间优化

	func findMaxForm(strs []string, m, n int) int {
	    f := make([][]int, m+1)
	    for i := range f {
	        f[i] = make([]int, n+1)
	    }
	    for _, s := range strs {
	        cnt0 := strings.Count(s, "0")
	        cnt1 := len(s) - cnt0
	        for j := m; j >= cnt0; j-- {
	            for k := n; k >= cnt1; k-- {
	                f[j][k] = max(f[j][k], f[j-cnt0][k-cnt1]+1)
	            }
	        }
	    }
	    return f[m][n]
	}
*/
func main() {

}
