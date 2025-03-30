package main

/**
Alice 在给 Bob 用手机打字。数字到字母的 对应 如下图所示。



 为了 打出 一个字母，Alice 需要 按 对应字母 i 次，i 是该字母在这个按键上所处的位置。


 比方说，为了按出字母 's' ，Alice 需要按 '7' 四次。类似的， Alice 需要按 '5' 两次得到字母 'k' 。
 注意，数字 '0' 和 '1' 不映射到任何字母，所以 Alice 不 使用它们。


 但是，由于传输的错误，Bob 没有收到 Alice 打字的字母信息，反而收到了 按键的字符串信息 。


 比方说，Alice 发出的信息为 "bob" ，Bob 将收到字符串 "2266622" 。


 给你一个字符串 pressedKeys ，表示 Bob 收到的字符串，请你返回 Alice 总共可能发出多少种文字信息 。

 由于答案可能很大，将它对 10⁹ + 7 取余 后返回。



 示例 1：


输入：pressedKeys = "22233"
输出：8
解释：
Alice 可能发出的文字信息包括：
"aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae" 和 "ce" 。
由于总共有 8 种可能的信息，所以我们返回 8 。


 示例 2：


输入：pressedKeys = "222222222222222222222222222222222222"
输出：82876089
解释：
总共有 2082876103 种 Alice 可能发出的文字信息。
由于我们需要将答案对 10⁹ + 7 取余，所以我们返回 2082876103 % (10⁹ + 7) = 82876089 。




 提示：


 1 <= pressedKeys.length <= 10⁵
 pressedKeys 只包含数字 '2' 到 '9' 。


 👍 123 👎 0

*/

/*
题型: dp
题目: 统计打字方案数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	const mx = 100_001

	var f = [mx]int{1, 1, 2, 4}
	var g = f

	for i := 4; i < mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}

	ans, cnt := 1, 0
	for i, c := range pressedKeys {
		cnt++
		if i == len(pressedKeys)-1 || byte(c) != pressedKeys[i+1] {
			if c != '7' && c != '9' {
				ans = ans * f[cnt] % mod
			} else {
				ans = ans * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
    2/3/4/5/6/8的状态方程: f(i) = f(i-1) + f(i-2) + f(i-3) i>=4
    7/9的状态方程: f(i) = f(i-1) + f(i-2) + f(i-3) + f(i-4) i>=5
边界条件:
    f(1) = 1
    f(2) = 2
    f(3) = 4
*/

/*
const mod = 1_000_000_007
const max = 100_001

var f = [max]int{1,1,2,4}
var g = f

	func init()  {
	    for i:=4;i<max;i++ {
	        f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
	        g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	    }
	}

	func countTexts(pressedKeys string) int {
	    ans ,cnt := 1,0
	    for i,c := range pressedKeys {
	        cnt++
	        if i == len(pressedKeys) - 1 || byte(c) != pressedKeys[i+1] {
	            if c != '7' && c != '9' {
	                ans = ans * f[cnt] % mod
	            } else {
	                ans = ans * g[cnt] % mod
	            }
	            cnt = 0
	        }
	    }
	    return ans
	}

时间复杂度: o(n), 其中n是pressedKeys的长度
空间复杂度: o(1)
*/
func main() {

}
