package main

import "slices"

/**
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：


 只使用数字1到9
 每个数字 最多使用一次


 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。



 示例 1:


输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。

 示例 2:


输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。

 示例 3:


输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。




 提示:


 2 <= k <= 9
 1 <= n <= 60


 Related Topics 数组 回溯 👍 927 👎 0

*/

/*
题型: 回溯
题目: 组合总和 III
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k, n int) (ans [][]int) {
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, t int) {
		d := k - len(path)              // 还要选 d 个数
		if t < 0 || t > (i*2-d+1)*d/2 { // 剪枝
			return
		}
		if d == 0 { // 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		// 不选 i
		if i > d {
			dfs(i-1, t)
		}

		// 选 i
		path = append(path, i)
		dfs(i-1, t-i)
		path = path[:len(path)-1]
	}
	dfs(9, n)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
枚举选哪个
func combinationSum3(k, n int) (ans [][]int) {
    path := []int{}
    var dfs func(int, int)
    dfs = func(i, t int) {
        d := k - len(path) // 还要选 d 个数
        if t < 0 || t > (i*2-d+1)*d/2 { // 剪枝
            return
        }
        if d == 0 { // 找到一个合法组合
            ans = append(ans, slices.Clone(path))
            return
        }
        for j := i; j >= d; j-- {
            path = append(path, j)
            dfs(j-1, t-j)
            path = path[:len(path)-1]
        }
    }
    dfs(9, n)
    return
}

*/

/*
选或不选

	func combinationSum3(k, n int) (ans [][]int) {
	    path := []int{}
	    var dfs func(int, int)
	    dfs = func(i, t int) {
	        d := k - len(path) // 还要选 d 个数
	        if t < 0 || t > (i*2-d+1)*d/2 { // 剪枝
	            return
	        }
	        if d == 0 { // 找到一个合法组合
	            ans = append(ans, slices.Clone(path))
	            return
	        }

	        // 不选 i
	        if i > d {
	            dfs(i-1, t)
	        }

	        // 选 i
	        path = append(path, i)
	        dfs(i-1, t-i)
	        path = path[:len(path)-1]
	    }
	    dfs(9, n)
	    return
	}
*/
func main() {

}
