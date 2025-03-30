package main

/**
来自未来的体育科学家给你两个整数数组 energyDrinkA 和 energyDrinkB，数组长度都等于 n。这两个数组分别代表 A、B 两种不同能量饮料每
小时所能提供的强化能量。

 你需要每小时饮用一种能量饮料来 最大化 你的总强化能量。然而，如果从一种能量饮料切换到另一种，你需要等待一小时来梳理身体的能量体系（在那个小时里你将不会获得任
何强化能量）。

 返回在接下来的 n 小时内你能获得的 最大 总强化能量。

 注意 你可以选择从饮用任意一种能量饮料开始。



 示例 1：


 输入：energyDrinkA = [1,3,1], energyDrinkB = [3,1,1]


 输出：5

 解释：

 要想获得 5 点强化能量，需要选择只饮用能量饮料 A（或者只饮用 B）。

 示例 2：


 输入：energyDrinkA = [4,1,1], energyDrinkB = [1,1,3]


 输出：7

 解释：


 第一个小时饮用能量饮料 A。
 切换到能量饮料 B ，在第二个小时无法获得强化能量。
 第三个小时饮用能量饮料 B ，并获得强化能量。




 提示：


 n == energyDrinkA.length == energyDrinkB.length
 3 <= n <= 10⁵
 1 <= energyDrinkA[i], energyDrinkB[i] <= 10⁵


 Related Topics 数组 动态规划 👍 47 👎 0

*/

/*
题型: dp
题目: 超级饮料的最大强化能量
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxEnergyBoost(a, b []int) int64 {
	n := len(a)
	c := [2][]int{a, b}
	memo := make([][2]int64, n)
	var dfs func(int, int) int64
	dfs = func(i, j int) int64 {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p == 0 { // 首次计算
			*p = max(dfs(i-1, j), dfs(i-2, j^1)) + int64(c[j][i])
		}
		return *p
	}
	return max(dfs(n-1, 0), dfs(n-1, 1))
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
	dfs(i,j)=max(dfs(i−1,j),dfs(i−2,j⊕1))+c[j][i]
边界条件
	dfs(−2,j)=dfs(−1,j)=0
*/

/*
递推
func maxEnergyBoost(a, b []int) int64 {
	n := len(a)
	f := make([][2]int64, n+2)
	for i, x := range a {
		f[i+2][0] = max(f[i+1][0], f[i][1]) + int64(x)
		f[i+2][1] = max(f[i+1][1], f[i][0]) + int64(b[i])
	}
	return max(f[n+1][0], f[n+1][1])
}

*/

func main() {

}
