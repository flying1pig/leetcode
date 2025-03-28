package main

import (
	"container/heap"
	"slices"
	"strings"
)

/**
给你一个字符串 s 。s[i] 要么是小写英文字母，要么是问号 '?' 。

 对于长度为 m 且 只 含有小写英文字母的字符串 t ，我们定义函数 cost(i) 为下标 i 之前（也就是范围 [0, i - 1] 中）出现过与 t[
i] 相同 字符出现的次数。

 字符串 t 的 分数 为所有下标 i 的 cost(i) 之 和 。

 比方说，字符串 t = "aab" ：


 cost(0) = 0
 cost(1) = 1
 cost(2) = 0
 所以，字符串 "aab" 的分数为 0 + 1 + 0 = 1 。


 你的任务是用小写英文字母 替换 s 中 所有 问号，使 s 的 分数最小 。

 请你返回替换所有问号 '?' 之后且分数最小的字符串。如果有多个字符串的 分数最小 ，那么返回字典序最小的一个。



 示例 1：


 输入：s = "???"


 输出： "abc"

 解释：这个例子中，我们将 s 中的问号 '?' 替换得到 "abc" 。

 对于字符串 "abc" ，cost(0) = 0 ，cost(1) = 0 和 cost(2) = 0 。

 "abc" 的分数为 0 。

 其他修改 s 得到分数 0 的字符串为 "cba" ，"abz" 和 "hey" 。

 这些字符串中，我们返回字典序最小的。

 示例 2：


 输入： s = "a?a?"


 输出： "abac"

 解释：这个例子中，我们将 s 中的问号 '?' 替换得到 "abac" 。

 对于字符串 "abac" ，cost(0) = 0 ，cost(1) = 0 ，cost(2) = 1 和 cost(3) = 0 。

 "abac" 的分数为 1 。



 提示：


 1 <= s.length <= 10⁵
 s[i] 要么是小写英文字母，要么是 '?' 。


 Related Topics 贪心 哈希表 字符串 计数 排序 堆（优先队列） 👍 15 👎 0

*/

/*
题型: 堆
题目: 替换字符串中的问号使分数最小
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minimizeStringValue(s string) string {
	h := make(hp, 26)
	for i := byte(0); i < 26; i++ {
		h[i].c = 'a' + i
	}
	for _, b := range s {
		if b != '?' {
			h[b-'a'].f++
		}
	}
	heap.Init(&h)

	t := make([]byte, strings.Count(s, "?"))
	for i := range t {
		t[i] = h[0].c
		h[0].f++ // 出现次数加一
		heap.Fix(&h, 0)
	}
	slices.Sort(t) // 排序，因为要求字典序最小

	ans := []byte(s)
	j := 0
	for i, b := range ans {
		if b == '?' {
			ans[i] = t[j] // 填入字母
			j++
		}
	}
	return string(ans)
}

type pair struct {
	f int
	c byte
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.f < b.f || a.f == b.f && a.c < b.c }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
