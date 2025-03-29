package main

import (
	"strconv"
	"strings"
)

/**
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。


 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和
"192.168@1.1" 是 无效 IP 地址。


 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序
或删除 s 中的任何数字。你可以按 任何 顺序返回答案。



 示例 1：


输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]


 示例 2：


输入：s = "0000"
输出：["0.0.0.0"]


 示例 3：


输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]




 提示：


 1 <= s.length <= 20
 s 仅由数字组成


 Related Topics 字符串 回溯 👍 1499 👎 0

*/

/*
题型: 回溯
题目: 复原 IP 地址
*/

// leetcode submit region begin(Prohibit modification and deletion)
var result []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	result = []string{}
	track := make([]string, 0)
	backtrack(s, track, 1)
	return result
}

func backtrack(s string, track []string, key int) {
	// 结束条件
	// key为段数，一共4段，选完4段，同时字符串选完时结束
	if key == 5 {
		if len(s) == 0 {
			str := strings.Join(track, ".")
			result = append(result, str)
		}
		return
	}
	// 选择列表
	// 每一段最大选择3位
	for j := 1; j <= 3; j++ {
		if j <= len(s) {
			// 选1-3位数字
			v, err := strconv.Atoi(s[:j])
			if err == nil && v <= 255 {
				// 做选择
				//fmt.Printf("第 %d 段，选择 %d位 s: %s\n", key, j, s[:j])
				track = append(track, s[:j])
				str := s[j:]
				// 下一段选择
				backtrack(str, track, key+1)
				// 撤销选择
				track = track[:len(track)-1]
			}
			// 每一段只能为0，不能为01
			if v == 0 {
				break
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
