package main

/**
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

 实现词典类 WordDictionary ：


 WordDictionary() 初始化词典对象
 void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
 bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些 '.
' ，每个 . 都可以表示任何一个字母。




 示例：


输入：
["WordDictionary","addWord","addWord","addWord","search","search","search",
"search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True
wordDictionary.search("b.."); // 返回 True




 提示：


 1 <= word.length <= 25
 addWord 中的 word 由小写英文字母组成
 search 中的 word 由 '.' 或小写英文字母组成
 最多调用 10⁴ 次 addWord 和 search


 Related Topics 深度优先搜索 设计 字典树 字符串 👍 604 👎 0

*/

/*
题型: 前缀树
题目: 添加与搜索单词 - 数据结构设计
*/

// leetcode submit region begin(Prohibit modification and deletion)
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieRoot *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

func (d *WordDictionary) AddWord(word string) {
	d.trieRoot.Insert(word)
}

func (d *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, d.trieRoot)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
