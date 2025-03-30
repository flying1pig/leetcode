package main

import "strconv"

/**
现有一串神秘的密文 ciphertext，经调查，密文的特点和规则如下：


 密文由非负整数组成
 数字 0-25 分别对应字母 a-z


 请根据上述规则将密文 ciphertext 解密为字母，并返回共有多少种解密结果。





 示例 1：


输入：ciphertext = 216612
输出：6
解释：216612 解密后有 6 种不同的形式，分别是 "cbggbc"，"vggbc"，"vggm"，"cbggm"，"cqgbc" 和 "cqgm"



 提示：


 0 <= ciphertext < 2³¹




 Related Topics 字符串 动态规划 👍 621 👎 0

*/

/*
题型: dp
题目: 解密数字
*/

// leetcode submit region begin(Prohibit modification and deletion)
func crackNumber(ciphertext int) int {
	src := strconv.Itoa(ciphertext)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程
f(i)=f(i−1)+f(i−2)[i−1≥0,10≤x≤25]
*/

func main() {

}
