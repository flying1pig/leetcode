package main

import "slices"

/**
一个魔法师有许多不同的咒语。

 给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。

 已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1
 或者 power[i] + 2 的咒语。

 每个咒语最多只能被使用 一次 。

 请你返回这个魔法师可以达到的伤害值之和的 最大值 。



 示例 1：


 输入：power = [1,1,3,4]


 输出：6

 解释：

 可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。

 示例 2：


 输入：power = [7,1,6,6]


 输出：13

 解释：

 可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。



 提示：


 1 <= power.length <= 10⁵
 1 <= power[i] <= 10⁹


 Related Topics 数组 哈希表 双指针 二分查找 动态规划 计数 排序 👍 46 👎 0

*/

/*
题型: dp
题目: 施咒的最大总伤害
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	f := make([]int, n+1)
	j := 0
	for i, x := range a {
		for a[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: dfs(i)=max(dfs(i−1),dfs(j−1)+a[i]⋅cnt[a[i]])
边界条件: dfs(-1) = 0
*/

/*
记忆化搜索
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := a[i]
		j := i
		for j > 0 && a[j-1] >= x-2 {
			j--
		}
		*p = max(dfs(i-1), dfs(j-1)+x*cnt[x])
		return *p
	}
	return int64(dfs(n - 1))
}

*/

/*
递推:
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	f := make([]int, n+1)
	j := 0
	for i, x := range a {
		for a[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}

*/

func main() {

}
