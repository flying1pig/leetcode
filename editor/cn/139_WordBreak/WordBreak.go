package main

/**
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。



 示例 1：


输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。


 示例 2：


输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。


 示例 3：


输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false




 提示：


 1 <= s.length <= 300
 1 <= wordDict.length <= 1000
 1 <= wordDict[i].length <= 20
 s 和 wordDict[i] 仅由小写英文字母组成
 wordDict 中的所有字符串 互不相同


 Related Topics 字典树 记忆化搜索 数组 哈希表 字符串 动态规划 👍 2711 👎 0

*/

/*
题型: 前缀树+动态规划
题目: 单词拆分
*/

// leetcode submit region begin(Prohibit modification and deletion)

type dicNode struct {
	children [26]*dicNode
	end      bool
}

func wordBreak(s string, wordDict []string) bool {
	dic := new(dicNode)
	for _, v := range wordDict {
		dic.add(v)
	}
	return dic.find(s, 0, make([]bool, len(s)))
}

func (d *dicNode) add(s string) {
	cur := d
	for _, v := range s {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = new(dicNode)
		}
		cur = cur.children[v-'a']
	}
	cur.end = true
}

func (d *dicNode) find(s string, index int, failed []bool) bool {
	if len(s[index:]) == 0 {
		return true
	}
	if failed[index] {
		return false
	}
	cur := d
	for i, v := range s[index:] {
		if cur.children[v-'a'] == nil {
			return false
		}
		cur = cur.children[v-'a']
		if !cur.end {
			continue
		}
		if d.find(s, index+i+1, failed) {
			return true
		}
		failed[index+i+1] = true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
