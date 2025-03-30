package main

import "math"

/**
给定一个整数数组 arr 和一个整数 k ，通过重复 k 次来修改数组。

 例如，如果 arr = [1, 2] ，
 k = 3 ，那么修改后的数组将是 [1, 2, 1, 2, 1, 2] 。

 返回修改后的数组中的最大的子数组之和。注意，子数组长度可以是 0，在这种情况下它的总和也是 0。

 由于 结果可能会很大，需要返回的
 10⁹ + 7 的 模 。



 示例 1：


输入：arr = [1,2], k = 3
输出：9


 示例 2：


输入：arr = [1,-2,1], k = 5
输出：2


 示例 3：


输入：arr = [-1,-2], k = 7
输出：0




 提示：



 1 <= arr.length <= 10⁵
 1 <= k <= 10⁵
 -10⁴ <= arr[i] <= 10⁴


 👍 151 👎 0

*/

/*
题型: dp
题目: K 次串联后最大子数组之和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func kConcatenationMaxSum(arr []int, k int) int {
	mod := 1_000_000_007
	sum := 0
	for _, val := range arr {
		sum = (sum + val) % mod
	}

	if k >= 2 {
		arr = append(arr, arr...)
	}
	ans := math.MinInt
	f := 0
	for _, x := range arr {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	if k <= 2 {
		return max(ans%mod, 0)
	}
	tmp := (max(sum, 0)*(k-2)%mod + ans%mod) % mod
	return max(tmp, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

/*
提示: The answer is the maximum between, the answer for k=1, the sum of the whole array multiplied by k,
or the maximum suffix sum plus the maximum prefix sum plus (k-2) multiplied by the whole array sum for k > 1.

数组最大前缀和+最大后缀和=数组拼接后的最大连续子数组和
计算过程中不要随意取模
*/

func main() {

}
