package main

/**
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：


 nums1[i] == nums2[j]
 且绘制的直线不与任何其他连线（非水平线）相交。


 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

 以这种方法绘制线条，并返回可以绘制的最大连线数。



 示例 1：


输入：nums1 = [1,4,2], nums2 = [1,2,4]
输出：2
解释：可以画出两条不交叉的线，如上图所示。
但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。




 示例 2：



输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
输出：3



 示例 3：



输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
输出：2



 提示：


 1 <= nums1.length, nums2.length <= 500
 1 <= nums1[i], nums2[j] <= 2000




 Related Topics 数组 动态规划 👍 635 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxUncrossedLines(s, t []int) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if s[i] == t[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(n-1, m-1)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
定义dfs(i,j)表示s[0:i]与t[0:j]之间的最大连线
状态转移方程:
	dfs(i,j) = dfs(i-1,j-1)+1 	s[i] = t[j]
	dfs(i,j) = max(dfs(i-1,j),dfs(i,j-1)) 	s[i] != t[j]
边界条件:
	dfs(-1,j) = dfs(i,-1) = 0
*/

/*
递归:
func maxUncrossedLines(s, t []int) int {
    n, m := len(s), len(t)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if i < 0 || j < 0 {
            return
        }
        p := &memo[i][j]
        if *p != -1 { // 之前计算过
            return *p
        }
        defer func() { *p = res }() // 记忆化
        if s[i] == t[j] {
            return dfs(i-1, j-1) + 1
        }
        return max(dfs(i-1, j), dfs(i, j-1))
    }
    return dfs(n-1, m-1)
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
递推:
func maxUncrossedLines(s, t []int) int {
    n, m := len(s), len(t)
    f := make([][]int, n+1)
    for i := range f {
        f[i] = make([]int, m+1)
    }
    for i, x := range s {
        for j, y := range t {
            if x == y {
                f[i+1][j+1] = f[i][j] + 1
            } else {
                f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
            }
        }
    }
    return f[n][m]
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
空间优化:
func maxUncrossedLines(s, t []int) int {
    m := len(t)
    f := make([]int, m+1)
    for _, x := range s {
        pre := 0 // f[0]
        for j, y := range t {
            if x == y {
                f[j+1], pre = pre+1, f[j+1]
            } else {
                pre = f[j+1]
                f[j+1] = max(f[j+1], f[j])
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
