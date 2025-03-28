package main

/**
单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：


 words.length == indices.length
 助记字符串 s 以 '#' 字符结尾
 对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与
words[i] 相等


 给你一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。



 示例 1：


输入：words = ["time", "me", "bell"]
输出：10
解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"


 示例 2：


输入：words = ["t"]
输出：2
解释：一组有效编码为 s = "t#" 和 indices = [0] 。




 提示：


 1 <= words.length <= 2000
 1 <= words[i].length <= 7
 words[i] 仅由小写字母组成


 Related Topics 字典树 数组 哈希表 字符串 👍 337 👎 0

*/

/*
题型: 前缀树
题目: 单词的压缩编码
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minimumLengthEncoding(words []string) int {
	head := &Node{}
	var root *Node
	var ans int
	var add bool
	for _, word := range words {
		root = head
		add = false
		for i := len(word) - 1; i >= 0; i-- {
			k := word[i] - 'a'
			if root.data[k] == nil {
				root.data[k] = &Node{}
				add = true
			}
			root = root.data[k]
			if root.data[26] != nil {
				ans -= len(word) - i + 1
				root.data[26] = nil
				add = true
			}

		}
		if add {
			ans += len(word) + 1
			root.data[26] = &Node{}
		}
	}
	return ans
}

type Node struct {
	data [27]*Node
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
