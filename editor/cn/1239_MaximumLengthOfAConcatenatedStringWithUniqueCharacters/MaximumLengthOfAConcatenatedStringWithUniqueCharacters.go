package main

import "math/bits"

/**
给定一个字符串数组 arr，字符串 s 是将 arr 的含有 不同字母 的 子序列 字符串 连接 所得的字符串。

 请返回所有可行解 s 中最长长度。

 子序列 是一种可以从另一个数组派生而来的数组，通过删除某些元素或不删除元素而不改变其余元素的顺序。



 示例 1：


输入：arr = ["un","iq","ue"]
输出：4
解释：所有可能的串联组合是：
- ""
- "un"
- "iq"
- "ue"
- "uniq" ("un" + "iq")
- "ique" ("iq" + "ue")
最大长度为 4。


 示例 2：


输入：arr = ["cha","r","act","ers"]
输出：6
解释：可能的解答有 "chaers" 和 "acters"。


 示例 3：


输入：arr = ["abcdefghijklmnopqrstuvwxyz"]
输出：26




 提示：


 1 <= arr.length <= 16
 1 <= arr[i].length <= 26
 arr[i] 中只含有小写英文字母


 Related Topics 位运算 数组 字符串 回溯 👍 246 👎 0

*/

/*
题型: 回溯
题目: 串联字符串的最大长度
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxLength(arr []string) (ans int) {
	masks := []int{}
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask>>ch&1 > 0 { // 若 mask 已有 ch，则说明 s 含有重复字母，无法构成可行解
				continue outer
			}
			mask |= 1 << ch // 将 ch 加入 mask 中
		}
		masks = append(masks, mask)
	}

	var backtrack func(int, int)
	backtrack = func(pos, mask int) {
		if pos == len(masks) {
			ans = max(ans, bits.OnesCount(uint(mask)))
			return
		}
		if mask&masks[pos] == 0 { // mask 和 masks[pos] 无公共元素
			backtrack(pos+1, mask|masks[pos])
		}
		backtrack(pos+1, mask)
	}
	backtrack(0, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
