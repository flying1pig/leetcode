package main

import (
	"sort"
	"strconv"
)

/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。



 示例 1：


输入：nums = [10,2]
输出："210"

 示例 2：


输入：nums = [3,30,34,5,9]
输出："9534330"




 提示：


 1 <= nums.length <= 100
 0 <= nums[i] <= 10⁹


 Related Topics 贪心 数组 字符串 排序 👍 1325 👎 0

*/

/*
题型: 贪心
题目: 最大数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool { //将数字转化成字符串，按数字第一个数字的大小递减排序，如果相等，则接着往下找
		s1, s2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		tmp1, tmp2 := s1+s2, s2+s1
		for k := 0; k < len(tmp1); k++ {
			if tmp1[k] == tmp2[k] {
				continue
			}
			return tmp1[k] > tmp2[k]
		}
		return true
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := ""
	for i := 0; i < len(nums); i++ {
		ans += strconv.Itoa(nums[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
