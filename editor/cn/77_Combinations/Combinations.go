package main

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

 你可以按 任何顺序 返回答案。



 示例 1：


输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

 示例 2：


输入：n = 1, k = 1
输出：[[1]]



 提示：


 1 <= n <= 20
 1 <= k <= n


 Related Topics 回溯 👍 1750 👎 0

*/

/*
题型: 回溯
题目: 组合
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n, k int) (ans [][]int) {
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path) // 还要选 d 个数
		if d == 0 {        // 选好了
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(n)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
枚举选哪个
func combine(n, k int) (ans [][]int) {
    path := []int{}
    var dfs func(int)
    dfs = func(i int) {
        d := k - len(path) // 还要选 d 个数
        if d == 0 { // 选好了
            ans = append(ans, append([]int(nil), path...))
            return
        }
        for j := i; j >= d; j-- {
            path = append(path, j)
            dfs(j - 1)
            path = path[:len(path)-1] // 恢复现场
        }
    }
    dfs(n)
    return
}

*/

/*
选或不选

	func combine(n, k int) (ans [][]int) {
	    path := []int{}
	    var dfs func(int)
	    dfs = func(i int) {
	        d := k - len(path) // 还要选 d 个数
	        if d == 0 { // 选好了
	            ans = append(ans, append([]int(nil), path...))
	            return
	        }

	        // 如果 i > d，可以不选 i
	        if i > d {
	            dfs(i - 1)
	        }

	        // 选 i
	        path = append(path, i)
	        dfs(i - 1)
	        path = path[:len(path)-1] // 恢复现场
	    }
	    dfs(n)
	    return
	}
*/
func main() {

}
