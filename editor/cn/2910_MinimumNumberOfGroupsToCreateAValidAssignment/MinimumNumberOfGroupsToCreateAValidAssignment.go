package main

/**
给你一组带编号的 balls 并要求将它们分类到盒子里，以便均衡地分配。你必须遵守两条规则：


 同一个盒子里的球必须具有相同的编号。但是，如果你有多个相同编号的球，你可以把它们放在不同的盒子里。
 最大的盒子只能比最小的盒子多一个球。


 返回遵循上述规则排列这些球所需要的盒子的最小数目。



 示例 1：


输入：balls = [3,2,3,2,3]
输出：2
解释：一个得到 2 个分组的方案如下，中括号内的数字都是下标：
我们可以如下排列 balls 到盒子里：
- [3,3,3]
- [2,2]
两个盒子之间的大小差没有超过 1。

 示例 2：


输入：balls = [10,10,10,3,1,1]
输出：4
解释：我们可以如下排列 balls 到盒子里：
- [10]
- [10,10]
- [3]
- [1,1]
无法得到一个遵循上述规则且小于 4 盒的答案。例如，把所有三个编号为 10 的球都放在一个盒子里，就会打破盒子之间最大尺寸差异的规则。




 提示：


 1 <= balls.length <= 10⁵
 1 <= balls[i] <= 10⁹


 Related Topics 贪心 数组 哈希表 👍 30 👎 0

*/

/*
题型: 贪心
题目: 合法分组的最少组数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, c := range cnt {
		k = min(k, c)
	}
	for ; ; k-- {
		ans := 0
		for _, c := range cnt {
			if c/k < c%k {
				ans = 0
				break
			}
			ans += (c + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
