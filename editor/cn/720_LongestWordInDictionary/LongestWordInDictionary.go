package main

/**
给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。

 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。

 请注意，单词应该从左到右构建，每个额外的字符都添加到前一个单词的结尾。



 示例 1：


输入：words = ["w","wo","wor","worl", "world"]
输出："world"
解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。


 示例 2：


输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出："apple"
解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"




 提示：


 1 <= words.length <= 1000
 1 <= words[i].length <= 30
 所有输入的字符串 words[i] 都只包含小写字母。


 Related Topics 字典树 数组 哈希表 字符串 排序 👍 359 👎 0

*/

/*
题型: 前缀树
题目: 词典中最长的单词
*/

// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func longestWord(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return longest
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
