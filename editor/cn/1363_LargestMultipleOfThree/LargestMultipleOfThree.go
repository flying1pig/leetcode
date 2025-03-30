package main

/**
给你一个整数数组 digits，你可以通过按 任意顺序 连接其中某些数字来形成 3 的倍数，请你返回所能得到的最大的 3 的倍数。

 由于答案可能不在整数数据类型范围内，请以字符串形式返回答案。如果无法得到答案，请返回一个空字符串。返回的结果不应包含不必要的前导零。



 示例 1：


输入：digits = [8,1,9]
输出："981"


 示例 2：


输入：digits = [8,6,7,1,0]
输出："8760"


 示例 3：


输入：digits = [1]
输出：""


 示例 4：


输入：digits = [0,0,0,0,0,0]
输出："0"




 提示：


 1 <= digits.length <= 10^4
 0 <= digits[i] <= 9


 Related Topics 贪心 数组 数学 动态规划 排序 👍 95 👎 0

*/

/*
题型: 字符串贪心
题目: 形成三的最大倍数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestMultipleOfThree(digits []int) string {
	/**
	思路：
	1.先排序
	2.计算总值
	3.判断总值能不能整除3，不能则删掉最小的
	*/

	record := make([]int, 10)
	mod := 0
	for _, v := range digits {
		mod = (v + mod) % 3
		record[v]++
	}

	var dfs func(mod int) int

	dfs = func(mod int) int {
		if mod == 0 {
			return 0
		}
		for i := mod; i < 9; i += 3 {
			if record[i] > 0 {
				record[i]--
				return i
			}
		}
		num1, num2 := -1, -1
		if mod == 1 {
			num1, num2 = dfs(2), dfs(2)
		} else {
			num1, num2 = dfs(1), dfs(1)
		}

		if num1 != -1 && num2 != -1 {
			return num1
		} else if num1 != -1 {
			record[num1]++
		} else if num2 != -1 {
			record[num2]++
		}
		return -1
	}

	if dfs(mod) == -1 {
		return ""
	}

	res := make([]byte, 0, len(digits))
	for i := 9; i >= 0; i-- {
		b := byte('0' + i)
		for j := 0; j < record[i]; j++ {
			if i == 0 && len(res) == 0 {
				return "0"
			}
			res = append(res, b)
		}
	}

	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
