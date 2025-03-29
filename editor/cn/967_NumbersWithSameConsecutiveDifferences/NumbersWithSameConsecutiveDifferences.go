package main

/**
返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。

 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。

 你可以按 任何顺序 返回答案。



 示例 1：


输入：n = 3, k = 7
输出：[181,292,707,818,929]
解释：注意，070 不是一个有效的数字，因为它有前导零。


 示例 2：


输入：n = 2, k = 1
输出：[10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]

 示例 3：


输入：n = 2, k = 0
输出：[11,22,33,44,55,66,77,88,99]


 示例 4：


输入：n = 2, k = 2
输出：[13,20,24,31,35,42,46,53,57,64,68,75,79,86,97]




 提示：


 2 <= n <= 9
 0 <= k <= 9


 Related Topics 广度优先搜索 回溯 👍 110 👎 0

*/

/*
题型: 回溯
题目: 连续差相同的数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numsSameConsecDiff(n int, k int) []int {
	var result []int = []int{}
	var dfs func(level, cur int)
	dfs = func(level, cur int) {
		// fmt.Printf("level=%d,cur=%d\n",level,cur)
		if level == n-1 {
			result = append(result, cur)
			return
		}
		//排除最高位是0的节点，剪枝。
		if cur*10 == cur {
			return
		}

		prev := cur % 10
		if prev >= k {
			dfs(level+1, cur*10+prev-k)
		}
		if k > 0 && prev+k < 10 {
			dfs(level+1, cur*10+prev+k)
		}
	}
	for i := 0; i < 10; i++ {
		dfs(0, i)
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
