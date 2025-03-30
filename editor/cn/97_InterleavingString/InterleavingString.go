package main

/**
给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：


 s = s1 + s2 + ... + sn
 t = t1 + t2 + ... + tm
 |n - m| <= 1
 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...


 注意：a + b 意味着字符串 a 和 b 连接。



 示例 1：


输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true


 示例 2：


输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false


 示例 3：


输入：s1 = "", s2 = "", s3 = ""
输出：true




 提示：


 0 <= s1.length, s2.length <= 100
 0 <= s3.length <= 200
 s1、s2、和 s3 都由小写英文字母组成




 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?

 Related Topics 字符串 动态规划 👍 1101 👎 0

*/

/*
题型: dp
题目: 交错字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1, s2, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		p := &memo[i+1][j+1]
		if *p < 0 { // 没有计算过
			if i >= 0 && s1[i] == s3[i+j+1] && dfs(i-1, j) ||
				j >= 0 && s2[j] == s3[i+j+1] && dfs(i, j-1) {
				*p = 1
			} else {
				*p = 0
			}
		}
		return *p == 1
	}
	return dfs(n-1, m-1)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	dfs(i,j) = (i>=0 && s1[i]=s3[i+j+1] && dfs(i-1,j)) || (j>0 && s2[j]=s3[i+j+1] && dfs(i,j-1))
边界条件
	dfs(-1,-1) = true
*/

/*
记忆化搜索
func isInterleave(s1, s2, s3 string) bool {
    n, m := len(s1), len(s2)
    if n+m != len(s3) {
        return false
    }

    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, m+1)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i < 0 && j < 0 {
            return true
        }
        p := &memo[i+1][j+1]
        if *p < 0 { // 没有计算过
            if i >= 0 && s1[i] == s3[i+j+1] && dfs(i-1, j) ||
               j >= 0 && s2[j] == s3[i+j+1] && dfs(i, j-1) {
                *p = 1
            } else {
                *p = 0
            }
        }
        return *p == 1
    }
    return dfs(n-1, m-1)
}

*/

/*
递推
func isInterleave(s1, s2, s3 string) bool {
    n, m := len(s1), len(s2)
    if n+m != len(s3) {
        return false
    }

    f := make([][]bool, n+1)
    for i := range f {
        f[i] = make([]bool, m+1)
    }
    f[0][0] = true
    for j := range m {
        f[0][j+1] = s2[j] == s3[j] && f[0][j]
    }
    for i := range n {
        f[i+1][0] = s1[i] == s3[i] && f[i][0]
        for j := range m {
            f[i+1][j+1] = s1[i] == s3[i+j+1] && f[i][j+1] ||
                          s2[j] == s3[i+j+1] && f[i+1][j]
        }
    }
    return f[n][m]
}

*/

/*
空间优化

	func isInterleave(s1, s2, s3 string) bool {
	    n, m := len(s1), len(s2)
	    if n+m != len(s3) {
	        return false
	    }

	    f := make([]bool, m+1)
	    f[0] = true
	    for j := range m {
	        f[j+1] = f[j] && s2[j] == s3[j]
	    }
	    for i := range n {
	        f[0] = f[0] && s1[i] == s3[i]
	        for j := range m {
	            f[j+1] = f[j+1] && s1[i] == s3[i+j+1] ||
	                     f[j] && s2[j] == s3[i+j+1]
	        }
	    }
	    return f[m]
	}
*/
func main() {

}
