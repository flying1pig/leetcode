package main

import (
	"container/heap"
	"sort"
)

/**
给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。

 返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。



 示例 1:


输入: s = "aab"
输出: "aba"


 示例 2:


输入: s = "aaab"
输出: ""




 提示:


 1 <= s.length <= 500
 s 只包含小写字母


 Related Topics 贪心 哈希表 字符串 计数 排序 堆（优先队列） 👍 554 👎 0

*/

/*
题型: 堆+贪心
题目: 重构字符串
*/

// leetcode submit region begin(Prohibit modification and deletion)
var cnt [26]int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return cnt[h.IntSlice[i]] > cnt[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cnt = [26]int{}
	maxCnt := 0
	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, c := range cnt[:] {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	ans := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if cnt[i]--; cnt[i] > 0 {
			h.push(i)
		}
		if cnt[j]--; cnt[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
