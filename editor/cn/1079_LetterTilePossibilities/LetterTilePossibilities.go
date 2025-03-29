package main

/**
你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。

 注意：本题中，每个活字字模只能使用一次。



 示例 1：


输入："AAB"
输出：8
解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。


 示例 2：


输入："AAABBC"
输出：188


 示例 3：


输入："V"
输出：1



 提示：


 1 <= tiles.length <= 7
 tiles 由大写英文字母组成


 Related Topics 哈希表 字符串 回溯 计数 👍 284 👎 0

*/

/*
题型: 回溯
题目: 活字印刷
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numTilePossibilities(tiles string) (ans int) {
	n := len(tiles)
	vis := make([]bool, n)
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			return
		}
		// 在当前层不能选择同样的字母，所以用哈希表来记录已经选过的字母
		// 用数组同理，因为该题字符集仅为所有大写英文字母
		m := make(map[byte]bool)
		for i := 0; i < n; i++ {
			if vis[i] || m[tiles[i]] {
				continue
			}
			vis[i], m[tiles[i]] = true, true
			ans++
			dfs(pos + 1)
			vis[i] = false
		}
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
