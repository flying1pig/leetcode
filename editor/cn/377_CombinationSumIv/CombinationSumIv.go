package main

/**
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

 题目数据保证答案符合 32 位整数范围。



 示例 1：


输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。


 示例 2：


输入：nums = [9], target = 3
输出：0




 提示：


 1 <= nums.length <= 200
 1 <= nums[i] <= 1000
 nums 中的所有元素 互不相同
 1 <= target <= 1000




 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

 👍 1114 👎 0

*/

/*
题型: dp
题目: 组合总和 Ⅳ
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum4(nums []int, target int) int {
	mem := make([]int, target+1)
	mem[0] = 1
	for i := 1; i <= target; i++ {
		for _, val := range nums {
			if val <= i {
				mem[i] += mem[i-val]
			}
		}
	}
	return mem[target]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: dfs(x) = dfs(x-nums[0]) + dfs(x-num[1]) + ... + dfs(x-num[n-1]), n = len(nums)
边界条件: dfs(0) = 1;
         dfs(i) = 0, i < 0;
*/

/*
递归:
func combinationSum4(nums []int, target int) int {
    var dfs func(int) int
    dfs = func(i int) int {
        ans := 0
        if i == 0 {
            return 1
        }
        if i < 0 {
            return 0
        }
        for _,val := range nums {
            ans += dfs(i-val)
        }
        return ans
    }
    return dfs(target)
}
时间复杂度: o(n^target)
空间复杂度: o(1)
*/

/*
记忆化搜索:
func combinationSum4(nums []int, target int) int {
    var dfs func(int) int
    mem := make([]int,target+1)
    for i := range mem {
        mem[i] = -1
    }
    dfs = func(i int) int {
        ans := 0
        if i == 0 {
            return 1
        }
        if i < 0 {
            return 0
        }
        if mem[i] != -1 {
            return mem[i]
        }
        for _,val := range nums {
            ans += dfs(i-val)
        }
        mem[i] = ans
        return ans
    }
    return dfs(target)
}
时间复杂度: o(target*n)
空间复杂度: o(target)
*/

/*
递推:
func combinationSum4(nums []int, target int) int {
    mem := make([]int,target+1)
    mem[0] = 1
    for i :=1;i<=target; i++ {
        for _,val := range nums {
            if val <=i {
                mem[i] += mem[i-val]
            }
        }
    }
    return mem[target]
}
时间复杂度: o(target*n)
空间复杂度: o(target)
*/

func main() {

}
