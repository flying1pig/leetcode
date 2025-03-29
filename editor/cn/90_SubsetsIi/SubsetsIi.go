package main

import "slices"

/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。

 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。







 示例 1：


输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]


 示例 2：


输入：nums = [0]
输出：[[],[0]]




 提示：


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10


 Related Topics 位运算 数组 回溯 👍 1301 👎 0

*/

/*
题型: 回溯
题目: 子集 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 选 x
		x := nums[i]
		path = append(path, x)
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复现场

		// 不选 x，跳过所有等于 x 的数
		// 如果不跳过这些数，会导致「选 x 不选 x'」和「不选 x 选 x'」这两种情况都会加到 ans 中，这就重复了
		i++
		for i < n && nums[i] == x {
			i++
		}
		dfs(i)
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
选或不选
func subsetsWithDup(nums []int) (ans [][]int) {
    slices.Sort(nums)
    n := len(nums)
    path := []int{}
    var dfs func(int)
    dfs = func(i int) {
        if i == n {
            ans = append(ans, slices.Clone(path))
            return
        }

        // 选 x
        x := nums[i]
        path = append(path, x)
        dfs(i + 1)
        path = path[:len(path)-1] // 恢复现场

        // 不选 x，跳过所有等于 x 的数
        // 如果不跳过这些数，会导致「选 x 不选 x'」和「不选 x 选 x'」这两种情况都会加到 ans 中，这就重复了
        i++
        for i < n && nums[i] == x {
            i++
        }
        dfs(i)
    }
    dfs(0)
    return ans
}

*/

/*
枚举选哪个

	func subsetsWithDup(nums []int) (ans [][]int) {
	    slices.Sort(nums)
	    n := len(nums)
	    path := []int{}
	    var dfs func(int)
	    dfs = func(i int) {
	        ans = append(ans, slices.Clone(path))

	        // 在 [i,n-1] 中选一个 nums[j]
	        // 注意选 nums[j] 意味着 [i,j-1] 中的数都没有选
	        for j := i; j < n; j++ {
	            // 如果 j>i，说明 nums[j-1] 没有选
	            // 同方法一，所有等于 nums[j-1] 的数都不选
	            if j > i && nums[j] == nums[j-1] {
	                continue
	            }
	            path = append(path, nums[j])
	            dfs(j + 1)
	            path = path[:len(path)-1] // 恢复现场
	        }
	    }
	    dfs(0)
	    return ans
	}
*/
func main() {

}
