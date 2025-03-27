package main

/**
给你字符串 s 和整数 k 。

 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

 英文中的 元音字母 为（a, e, i, o, u）。



 示例 1：

 输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母。


 示例 2：

 输入：s = "aeiou", k = 2
输出：2
解释：任意长度为 2 的子字符串都包含 2 个元音字母。


 示例 3：

 输入：s = "leetcode", k = 3
输出：2
解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。


 示例 4：

 输入：s = "rhythms", k = 4
输出：0
解释：字符串 s 中不含任何元音字母。


 示例 5：

 输入：s = "tryhard", k = 4
输出：1




 提示：


 1 <= s.length <= 10^5
 s 由小写英文字母组成
 1 <= k <= s.length


 Related Topics 字符串 滑动窗口 👍 159 👎 0

*/

/*
题型: 定长滑动窗口
题目: 定长子串中元音的最大数目
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxVowels(s string, k int) (ans int) {
	vowel := 0
	for i, in := range s {
		// 1. 进入窗口
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		ans = max(ans, vowel)
		// 3. 离开窗口
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return
}

//时间复杂度: o(n)
//空间复杂度: o(1)
//leetcode submit region end(Prohibit modification and deletion)

/*
定长滑动窗口套路:
1、入：下标为 i 的元素进入窗口，更新相关统计量。如果 i<k−1 则重复第一步。
2、更新：更新答案。一般是更新最大值/最小值。
3、出：下标为 i−k+1 的元素离开窗口，更新相关统计量。
*/
func main() {

}

/*
定长滑动窗口选做题目:
3439. 重新安排会议得到最多空余时间 I 1729
2134. 最少交换次数来组合所有的 1 II 1748
1297. 子串的最大出现次数 1748
2653. 滑动子数组的美丽值 1786
1888. 使二进制字符串字符交替的最少反转次数 2006
567. 字符串的排列
438. 找到字符串中所有字母异位词
30. 串联所有单词的子串
2156. 查找给定哈希值的子串 2063
2953. 统计完全子字符串 2449
1016. 子串能表示从 1 到 N 数字的二进制串 做到 O(∣s∣)
683. K 个关闭的灯泡（会员题）做到 O(n)
2067. 等计数子串的数量（会员题）
2524. 子数组的最大频率分数（会员题）
*/
