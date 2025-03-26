package main

/**
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。

 你可以对一个单词进行如下三种操作：


 插入一个字符
 删除一个字符
 替换一个字符




 示例 1：


输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')


 示例 2：


输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')




 提示：


 0 <= word1.length, word2.length <= 500
 word1 和 word2 由小写英文字母组成


 Related Topics 字符串 动态规划 👍 3638 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++ // f[0] = i + 1
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(f[j+1], f[j], pre)+1, f[j+1]
			}
		}
	}
	return f[m]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
用dp[i][j]表示A的前i个字母和B的前j个字母之间的编辑距离。
	对单词A删除一个字符和对单词B插入一个字符是等价的
	对单词A插入一个字符和对单词B删除一个字符是等价的
	对单词A替换一个字符和对单词B替换一个字符是等价的
所以，对单词A的操作可以转换为:
	在单词A中插入一个字符
	在单词B中插入一个字符
	修改单词A的一个字符

这样，我们就可以把原问题转换为规模较小的子问题:
	在单词 A 中插入一个字符：如果我们知道 horse 到 ro 的编辑距离为 a，
那么显然 horse 到 ros 的编辑距离不会超过 a + 1。这是因为我们可以在 a
次操作后将 horse 和 ro 变为相同的字符串，只需要额外的 1 次操作，在单词 A 的末尾添加字符 s，
就能在 a + 1 次操作后将 horse 和 ro 变为相同的字符串;
	在单词 B 中插入一个字符：如果我们知道 hors 到 ros 的编辑距离为 b，那么显然 horse 到
ros 的编辑距离不会超过 b + 1，原因同上；
	修改单词 A 的一个字符：如果我们知道 hors 到 ro 的编辑距离为 c，那么显然 horse 到
ros 的编辑距离不会超过 c + 1，原因同上。

所以，当我们获得dp[i][j-1],dp[i-1][j],dp[i-1][j-1]的值后就可以计算出dp[i][j]
	dp[i][j-1]为 A 的前 i 个字符和 B 的前 j - 1 个字符编辑距离的子问题。即对于 B
的第 j 个字符，我们在 A 的末尾添加了一个相同的字符，那么 dp[i][j] 最小可以为 dp[i][j-1] + 1;
	dp[i-1][j]为 A 的前 i - 1 个字符和 B 的前 j 个字符编辑距离的子问题。即对于 A
的第 i 个字符，我们在 B 的末尾添加了一个相同的字符，那么 dp[i][j] 最小可以为 dp[i-1][j] + 1;
	dp[i-1][j-1] 为 A 前 i - 1 个字符和 B 的前 j - 1 个字符编辑距离的子问题。即对于 B
的第 j 个字符，我们修改 A 的第 i 个字符使它们相同，那么 dp[i][j] 最小可以为 dp[i-1][j-1] + 1。
特别地，如果 A 的第 i 个字符和 B 的第 j 个字符原本就相同，那么我们实际上不需要进行修改操作。
在这种情况下，dp[i][j] 最小可以为 dp[i-1][j-1]。

状态方程:
若A和B的最后一个字母相同:
	dp[i][j] = dp[i-1][j-1]
若A和B的最后一个字符不同:
	dp[i][j] = 1+min(dp[i][j-1],dp[i-1][j],dp[i-1][j-1])

边界条件:
	dp[i][0] = i
	dp[0][j] = j
*/

/*
记忆化搜索:
func minDistance(s, t string) int {
    n, m := len(s), len(t)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示还没有计算过
        }
    }
    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if i < 0 {
            return j + 1
        }
        if j < 0 {
            return i + 1
        }
        p := &memo[i][j]
        if *p != -1 { // 之前计算过
            return *p
        }
        defer func() { *p = res }() // 记忆化
        if s[i] == t[j] {
            return dfs(i-1, j-1)
        }
        return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
    }
    return dfs(n-1, m-1)
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
递推:
func minDistance(s, t string) int {
    n, m := len(s), len(t)
    f := make([][]int, n+1)
    for i := range f {
        f[i] = make([]int, m+1)
    }
    for j := 1; j <= m; j++ {
        f[0][j] = j
    }
    for i, x := range s {
        f[i+1][0] = i + 1
        for j, y := range t {
            if x == y {
                f[i+1][j+1] = f[i][j]
            } else {
                f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
            }
        }
    }
    return f[n][m]
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
空间优化: 连个数组
func minDistance(s, t string) int {
    n, m := len(s), len(t)
    f := [2][]int{make([]int, m+1), make([]int, m+1)}
    for j := 1; j <= m; j++ {
        f[0][j] = j
    }
    for i, x := range s {
        f[(i+1)%2][0] = i + 1
        for j, y := range t {
            if x == y {
                f[(i+1)%2][j+1] = f[i%2][j]
            } else {
                f[(i+1)%2][j+1] = min(f[i%2][j+1], f[(i+1)%2][j], f[i%2][j]) + 1
            }
        }
    }
    return f[n%2][m]
}
时间复杂度: o(mn)
空间复杂度: o(m)
*/

/*
空间优化，一个数组:
func minDistance(s, t string) int {
    m := len(t)
    f := make([]int, m+1)
    for j := 1; j <= m; j++ {
        f[j] = j
    }
    for _, x := range s {
        pre := f[0]
        f[0]++ // f[0] = i + 1
        for j, y := range t {
            if x == y {
                f[j+1], pre = pre, f[j+1]
            } else {
                f[j+1], pre = min(f[j+1], f[j], pre)+1, f[j+1]
            }
        }
    }
    return f[m]
}
时间复杂度: o(mn)
空间复杂度: o(m)
*/

func main() {

}
