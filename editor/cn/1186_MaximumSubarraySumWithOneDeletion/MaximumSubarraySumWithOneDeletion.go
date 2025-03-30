package main

import "math"

/**
给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以决定
要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。

 注意，删除一个元素后，子数组 不能为空。



 示例 1：


输入：arr = [1,-2,0,3]
输出：4
解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。

 示例 2：


输入：arr = [1,-2,-2,3]
输出：3
解释：我们直接选出 [3]，这就是最大和。


 示例 3：


输入：arr = [-1,-1,-1,-1]
输出：-1
解释：最后得到的子数组不能为空，所以我们不能选择 [-1] 并从中删去 -1 来得到 0。
     我们应该直接选择 [-1]，或者选择 [-1, -1] 再从中删去一个 -1。




 提示：



 1 <= arr.length <= 10⁵
 -10⁴ <= arr[i] <= 10⁴


 Related Topics 数组 动态规划 👍 372 👎 0

*/

/*
题型: dp
题目: 删除一次得到子数组最大和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSum(arr []int) int {
	memo := make([][2]int, len(arr))
	for i := range memo {
		memo[i] = [2]int{math.MinInt, math.MinInt}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return math.MinInt / 2 // 除 2 防止负数相加溢出
		}
		p := &memo[i][j]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if j == 0 {
			return max(dfs(i-1, 0), 0) + arr[i]
		}
		return max(dfs(i-1, 1)+arr[i], dfs(i-1, 0))
	}
	ans := math.MinInt
	for i := range arr {
		ans = max(ans, dfs(i, 0), dfs(i, 1))
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	dfs(i,0)=max(dfs(i−1,0),0)+arr[i]
	dfs(i,1)=max(dfs(i−1,1)+arr[i],dfs(i−1,0))

*/

/*
递推
func maximumSum(arr []int) int {
    ans := math.MinInt
    f := make([][2]int, len(arr)+1)
    f[0] = [2]int{math.MinInt / 2, math.MinInt / 2} // 除 2 防止负数相加溢出
    for i, x := range arr {
        f[i+1][0] = max(f[i][0], 0) + x
        f[i+1][1] = max(f[i][1]+x, f[i][0])
        ans = max(ans, f[i+1][0], f[i+1][1])
    }
    return ans
}

*/

/*
空间优化

	func maximumSum(arr []int) int {
	    ans := math.MinInt / 2
	    f0, f1 := ans, ans
	    for _, x := range arr {
	        f1 = max(f1+x, f0)
	        f0 = max(f0, 0) + x
	        ans = max(ans, f0, f1)
	    }
	    return ans
	}
*/
func main() {

}
