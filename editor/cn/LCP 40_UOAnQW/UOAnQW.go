package main

import "slices"

/**
「力扣挑战赛」心算项目的挑战比赛中，要求选手从 `N` 张卡牌中选出 `cnt` 张卡牌，若这 `cnt` 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 `
cnt` 张卡牌数字总和。
给定数组 `cards` 和 `cnt`，其中 `cards[i]` 表示第 `i` 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡
牌方案，则返回 0。

**示例 1：**

> 输入：`cards = [1,2,8,9], cnt = 3`
>
> 输出：`18`
>
> 解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。

**示例 2：**

> 输入：`cards = [3,3,1], cnt = 1`
>
> 输出：`0`
>
> 解释：不存在获取有效得分的卡牌方案。

**提示：**
- `1 <= cnt <= cards.length <= 10^5`
- `1 <= cards[i] <= 1000`

 Related Topics 贪心 数组 排序 👍 130 👎 0

*/

/*
题型: 贪心
题目: 心算挑战
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumScore(cards []int, cnt int) int {
	slices.SortFunc(cards, func(a, b int) int { return b - a })
	s := 0
	for _, v := range cards[:cnt] {
		s += v
	}
	if s%2 == 0 { // s 是偶数
		return s
	}

	replacedSum := func(x int) int {
		for _, v := range cards[cnt:] {
			if v%2 != x%2 { // 找到一个最大的奇偶性和 x 不同的数
				return s - x + v // 用 v 替换 s
			}
		}
		return 0
	}

	x := cards[cnt-1]
	ans := replacedSum(x)           // 替换 x
	for i := cnt - 2; i >= 0; i-- { // 前 cnt-1 个数
		if cards[i]%2 != x%2 { // 找到一个最小的奇偶性和 x 不同的数
			ans = max(ans, replacedSum(cards[i])) // 替换
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
