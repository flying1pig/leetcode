package main

import (
	"slices"
)

/**
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几
天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。



 示例 1:


输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]


 示例 2:


输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]


 示例 3:


输入: temperatures = [30,60,90]
输出: [1,1,0]



 提示：


 1 <= temperatures.length <= 10⁵
 30 <= temperatures[i] <= 100


 Related Topics 栈 数组 单调栈 👍 1930 👎 0

*/

/*
题型: 单调栈
题目: 每日温度
*/

// leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{}
	for i, t := range slices.Backward(temperatures) {
		for len(st) > 0 && t >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

//时间复杂度：O(n)，其中 n 为 temperatures 的长度。
//空间复杂度：O(min(n,U))，其中 U=max(temperatures)−min(temperatures)+1
//leetcode submit region end(Prohibit modification and deletion)

/*
从右到左
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{}
	for i, t := range slices.Backward(temperatures) {
		for len(st) > 0 && t >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

//时间复杂度：O(n)，其中 n 为 temperatures 的长度。
//空间复杂度：O(min(n,U))，其中 U=max(temperatures)−min(temperatures)+1
*/

/*
从左到右
func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    st := []int{} // todolist
    for i, t := range temperatures {
        for len(st) > 0 && t > temperatures[st[len(st)-1]] {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            ans[j] = i - j
        }
        st = append(st, i)
    }
    return ans
}
时间复杂度：O(n)，其中 n 为 temperatures 的长度。
空间复杂度：O(n)。
*/

func main() {
}
