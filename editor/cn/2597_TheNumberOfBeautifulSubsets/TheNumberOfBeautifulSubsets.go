package main

/**
给你一个由正整数组成的数组 nums 和一个 正 整数 k 。

 如果 nums 的子集中，任意两个整数的绝对差均不等于 k ，则认为该子数组是一个 美丽 子集。

 返回数组 nums 中 非空 且 美丽 的子集数目。

 nums 的子集定义为：可以经由 nums 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集
。



 示例 1：


输入：nums = [2,4,6], k = 2
输出：4
解释：数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
可以证明数组 [2,4,6] 中只存在 4 个美丽子集。


 示例 2：


输入：nums = [1], k = 1
输出：1
解释：数组 nums 中的美丽数组有：[1] 。
可以证明数组 [1] 中只存在 1 个美丽子集。




 提示：


 1 <= nums.length <= 18
 1 <= nums[i], k <= 1000


 Related Topics 数组 哈希表 数学 动态规划 回溯 组合数学 排序 👍 92 👎 0

*/

/*
题型: 回溯
题目: 美丽子集的数目
*/

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulSubsets(nums []int, k int) int {
	ans := -1 // 去掉空集
	cnt := map[int]int{}

	var dfs func(int)
	dfs = func(i int) {
		ans++
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ { // 枚举选哪个
			x := nums[j]
			if cnt[x-k] == 0 && cnt[x+k] == 0 { // 可以选
				cnt[x]++   // 选
				dfs(j + 1) // 下一个数在 [j+1, n-1] 中选
				cnt[x]--   // 撤销，恢复现场
			}
		}
	}

	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
输入视角，选或不选
func beautifulSubsets(nums []int, k int) int {
    ans := -1 // 去掉空集
    cnt := map[int]int{}

    // nums[i] 选或不选
    var dfs func(int)
    dfs = func(i int) {
        if i == len(nums) {
            ans++
            return
        }
        dfs(i + 1) // 不选
        x := nums[i]
        if cnt[x-k] == 0 && cnt[x+k] == 0 { // 可以选
            cnt[x]++ // 选
            dfs(i + 1) // 讨论 nums[i+1] 选或不选
            cnt[x]-- // 撤销，恢复现场
        }
    }

    dfs(0)
    return ans
}

*/

/*
答案视角，枚举选哪个

	func beautifulSubsets(nums []int, k int) int {
	    ans := -1 // 去掉空集
	    cnt := map[int]int{}

	    var dfs func(int)
	    dfs = func(i int) {
	        ans++
	        if i == len(nums) {
	            return
	        }
	        for j := i; j < len(nums); j++ { // 枚举选哪个
	            x := nums[j]
	            if cnt[x-k] == 0 && cnt[x+k] == 0 { // 可以选
	                cnt[x]++ // 选
	                dfs(j + 1) // 下一个数在 [j+1, n-1] 中选
	                cnt[x]-- // 撤销，恢复现场
	            }
	        }
	    }

	    dfs(0)
	    return ans
	}
*/
func main() {

}
