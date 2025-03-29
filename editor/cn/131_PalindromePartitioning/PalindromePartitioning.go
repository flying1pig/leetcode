package main

import "slices"

/**
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。



 示例 1：


输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]


 示例 2：


输入：s = "a"
输出：[["a"]]




 提示：


 1 <= s.length <= 16
 s 仅由小写英文字母组成


 Related Topics 字符串 动态规划 回溯 👍 2018 👎 0

*/

/*
题型: 回溯
题目: 分割回文串
*/

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) (ans [][]string) {
	n := len(s)
	path := []string{}

	// 考虑 i 后面的逗号怎么选
	// start 表示当前这段回文子串的开始位置
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n { // s 分割完毕
			ans = append(ans, slices.Clone(path))
			return
		}

		// 不选 i 和 i+1 之间的逗号（i=n-1 时一定要选）
		if i < n-1 {
			// 考虑 i+1 后面的逗号怎么选
			dfs(i+1, start)
		}

		// 选 i 和 i+1 之间的逗号（把 s[i] 作为子串的最后一个字符）
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			// 考虑 i+1 后面的逗号怎么选
			// start=i+1 表示下一个子串从 i+1 开始
			dfs(i+1, i+1)
			path = path[:len(path)-1] // 恢复现场
		}
	}

	dfs(0, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
输入视角，逗号选或不选
func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func partition(s string) (ans [][]string) {
    n := len(s)
    path := []string{}

    // 考虑 i 后面的逗号怎么选
    // start 表示当前这段回文子串的开始位置
    var dfs func(int, int)
    dfs = func(i, start int) {
        if i == n { // s 分割完毕
            ans = append(ans, slices.Clone(path))
            return
        }

        // 不选 i 和 i+1 之间的逗号（i=n-1 时一定要选）
        if i < n-1 {
            // 考虑 i+1 后面的逗号怎么选
            dfs(i+1, start)
        }

        // 选 i 和 i+1 之间的逗号（把 s[i] 作为子串的最后一个字符）
        if isPalindrome(s, start, i) {
            path = append(path, s[start:i+1])
            // 考虑 i+1 后面的逗号怎么选
            // start=i+1 表示下一个子串从 i+1 开始
            dfs(i+1, i+1)
            path = path[:len(path)-1] // 恢复现场
        }
    }

    dfs(0, 0)
    return
}

*/

/*
答案视角，枚举子串结束位置
func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func partition(s string) (ans [][]string) {
    n := len(s)
    path := []string{}

    // 考虑 s[i:] 怎么分割
    var dfs func(int)
    dfs = func(i int) {
        if i == n { // s 分割完毕
            ans = append(ans, slices.Clone(path))
            return
        }
        for j := i; j < n; j++ { // 枚举子串的结束位置
            if isPalindrome(s, i, j) {
                path = append(path, s[i:j+1]) // 分割！
                // 考虑剩余的 s[j+1:] 怎么分割
                dfs(j + 1)
                path = path[:len(path)-1] // 恢复现场
            }
        }
    }

    dfs(0)
    return
}

*/

func main() {

}
