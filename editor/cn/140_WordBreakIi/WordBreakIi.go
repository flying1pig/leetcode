package main

/**
给定一个字符串 s 和一个字符串字典
 wordDict ，在字符串
 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。

 注意：词典中的同一个单词可能在分段中被重复使用多次。



 示例 1：


输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
输出:["cats and dog","cat sand dog"]


 示例 2：


输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine",
"pineapple"]
输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
解释: 注意你可以重复使用字典中的单词。


 示例 3：


输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
输出:[]




 提示：





 1 <= s.length <= 20
 1 <= wordDict.length <= 1000
 1 <= wordDict[i].length <= 10
 s 和 wordDict[i] 仅有小写英文字母组成
 wordDict 中所有字符串都 不同


 Related Topics 字典树 记忆化搜索 数组 哈希表 字符串 动态规划 回溯 👍 781 👎 0

*/

/*
题型: 回溯
题目: 单词拆分 II
*/

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool, len(s)+1)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}

	ret := dealdata(s, wordDict, 0, m)
	return ret
}

func dealdata(s string, wordDict []string, start int, m map[string]bool) []string {
	var res []string
	if start == len(s) {
		res = append(res, "")
	}

	for end := start + 1; end <= len(s); end++ {
		// 判断s[start:end] 是否在 wordDict中
		if _, ok := m[s[start:end]]; ok {
			list := dealdata(s, wordDict, end, m)
			for _, v := range list {
				var tmp string
				if v == "" {
					tmp = ""
				} else {
					tmp = " "
				}
				res = append(res, s[start:end]+tmp+v)
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
