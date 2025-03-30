package main

/**
给你一个非负整数数组 nums 和一个整数 target 。

 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：


 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。


 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。



 示例 1：


输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3


 示例 2：


输入：nums = [1], target = 1
输出：1




 提示：


 1 <= nums.length <= 20
 0 <= nums[i] <= 1000
 0 <= sum(nums[i]) <= 1000
 -1000 <= target <= 1000


 Related Topics 数组 动态规划 回溯 👍 2109 👎 0

*/

/*
题型: dp
题目: 目标和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2 //背包容量

	n := len(nums)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, m+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var f func(int, int) int //第一个参数是待选数组长度，第二个参数是背包容量
	f = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if mem[i][c] != -1 {
			return mem[i][c]
		}
		if c < nums[i] { //不选
			mem[i][c] = f(i-1, c)
			return mem[i][c]
		}
		mem[i][c] = f(i-1, c) + f(i-1, c-nums[i])
		return mem[i][c] //选和不选
	}
	return f(n-1, m)
}

func abs(i int) int {
	if i < 0 {
		return i
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

/*
这个题要先做一下数学转换。
问题等价于从原数组中取几个数，使得和为m = (s- |target|)/2，如果m小于0或者不为整数，则原问题没得答案。
*/

/*
递归:
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2 //背包容量

	n := len(nums)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, m+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var f func(int, int) int //第一个参数是待选数组长度，第二个参数是背包容量
	f = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if mem[i][c] != -1 {
			return mem[i][c]
		}
		if c < nums[i] {	//不选
			return f(i-1,c)
		}
		return f(i-1,c) + f(i-1,c-nums[i])	//选和不选
	}
	return f(n-1,m)
}

func abs(i int) int {
	if i < 0 {
		return i
	}
	return i
}
时间复杂度: o(nm), n为数组长度，m为数组元素和减去target的绝对值。
空间复杂度: o(nm)
*/

/*
递推:
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2 //背包容量

	n := len(nums)
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, m+1)
	}

	mem[0][0] =1
	for i,x := range nums {
		for c :=0;c<=m;c++ {
			if c <x {
				mem[i+1][c] = mem[i][c]  //不选
			} else {
				mem[i+1][c] = mem[i][c] + mem[i][c-x]
			}
		}
	}
	return mem[n][m]
}

func abs(i int) int {
	if i < 0 {
		return i
	}
	return i
}
时间复杂度: o(mn)
空间复杂度: o(mn)
*/

/*
空间优化:
	func findTargetSumWays(nums []int, target int) int {
		s := 0
		for _, x := range nums {
			s += x
		}
		s -= abs(target)
		if s < 0 || s%2 == 1 {
			return 0
		}
		m := s / 2 //背包容量

		n := len(nums)
		mem := make([][]int, 2)
		for i := range mem {
			mem[i] = make([]int, m+1)
		}

		mem[0][0] = 1
		for i, x := range nums {
			for c := 0; c <= m; c++ {
				if c < x {
					mem[(i+1)%2][c] = mem[i%2][c] //不选
				} else {
					mem[(i+1)%2][c] = mem[i%2][c] + mem[i%2][c-x] //不选+选
				}
			}
		}
		return mem[n%2][m]
	}

	func abs(i int) int {
		if i < 0 {
			return i
		}
		return i
	}

时间复杂度: o(nm)
空间复杂度: o(m)
*/

/*
空间优化, 一个数组

	func findTargetSumWays(nums []int, target int) int {
	    s := 0
	    for _, x := range nums {
	        s += x
	    }
	    s -= abs(target)
	    if s < 0 || s%2 == 1 {
	        return 0
	    }
	    m := s / 2

	    f := make([]int, m+1)
	    f[0] = 1
	    for _, x := range nums {
	        for c := m; c >= x; c-- {
	            f[c] += f[c-x]
	        }
	    }
	    return f[m]
	}

func abs(x int) int { if x < 0 { return -x }; return x }
时间复杂度: o(nm)
空间复杂度: o(m)
*/
func main() {

}
