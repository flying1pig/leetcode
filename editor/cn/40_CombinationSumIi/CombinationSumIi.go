package main

import "slices"

/**
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

 candidates 中的每个数字在每个组合中只能使用 一次 。

 注意：解集不能包含重复的组合。



 示例 1:


输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

 示例 2:


输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]



 提示:


 1 <= candidates.length <= 100
 1 <= candidates[i] <= 50
 1 <= target <= 30


 Related Topics 数组 回溯 👍 1665 👎 0

*/

/*
题型: 回溯
题目: 组合总和 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	n := len(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, left int) {
		// 所选元素之和恰好等于 target
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 没有可以选的数字
		if i == n {
			return
		}

		// 所选元素之和无法恰好等于 target
		x := candidates[i]
		if left < x {
			return
		}

		// 选 x
		path = append(path, x)
		dfs(i+1, left-x)
		path = path[:len(path)-1] // 恢复现场

		// 不选 x，那么后面所有等于 x 的数都不选
		// 如果不跳过这些数，会导致「选 x 不选 x'」和「不选 x 选 x'」这两种情况都会加到 ans 中，这就重复了
		i++
		for i < n && candidates[i] == x {
			i++
		}
		dfs(i, left)
	}
	dfs(0, target)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
选或不选
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	n := len(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, left int) {
		// 所选元素之和恰好等于 target
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 没有可以选的数字
		if i == n {
			return
		}

		// 所选元素之和无法恰好等于 target
		x := candidates[i]
		if left < x {
			return
		}

		// 选 x
		path = append(path, x)
		dfs(i+1, left-x)
		path = path[:len(path)-1] // 恢复现场

		// 不选 x，那么后面所有等于 x 的数都不选
		// 如果不跳过这些数，会导致「选 x 不选 x'」和「不选 x 选 x'」这两种情况都会加到 ans 中，这就重复了
		i++
		for i < n && candidates[i] == x {
			i++
		}
		dfs(i, left)
	}
	dfs(0, target)
	return ans
}
*/

/*
枚举选哪个
func combinationSum2(candidates []int, target int) (ans [][]int) {
    slices.Sort(candidates)
    path := []int{}
    var dfs func(int, int)
    dfs = func(i, left int) {
        // 所选元素之和恰好等于 target
        if left == 0 {
            ans = append(ans, slices.Clone(path))
            return
        }

        // 在 [i, len(candidates)-1] 中选一个 candidates[j]
        // 注意选 candidates[j] 意味着 [i,j-1] 中的数都没有选
        for j := i; j < len(candidates) && candidates[j] <= left; j++ {
            // 如果 j>i，说明 candidates[j-1] 没有选
            // 同方法一，所有等于 candidates[j-1] 的数都不选
            if j > i && candidates[j] == candidates[j-1] {
                continue
            }
            path = append(path, candidates[j])
            dfs(j+1, left-candidates[j])
            path = path[:len(path)-1] // 恢复现场
        }
    }
    dfs(0, target)
    return ans
}

*/

func main() {

}
