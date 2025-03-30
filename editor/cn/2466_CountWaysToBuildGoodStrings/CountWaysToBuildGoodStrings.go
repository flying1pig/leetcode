package main

/**
给你整数 zero ，one ，low 和 high ，我们从空字符串开始构造一个字符串，每一步执行下面操作中的一种：


 将 '0' 在字符串末尾添加 zero 次。
 将 '1' 在字符串末尾添加 one 次。


 以上操作可以执行任意次。

 如果通过以上过程得到一个 长度 在 low 和 high 之间（包含上下边界）的字符串，那么这个字符串我们称为 好 字符串。

 请你返回满足以上要求的 不同 好字符串数目。由于答案可能很大，请将结果对 10⁹ + 7 取余 后返回。



 示例 1：

 输入：low = 3, high = 3, zero = 1, one = 1
输出：8
解释：
一个可能的好字符串是 "011" 。
可以这样构造得到："" -> "0" -> "01" -> "011" 。
从 "000" 到 "111" 之间所有的二进制字符串都是好字符串。


 示例 2：

 输入：low = 2, high = 3, zero = 1, one = 2
输出：5
解释：好字符串为 "00" ，"11" ，"000" ，"110" 和 "011" 。




 提示：


 1 <= low <= high <= 10⁵
 1 <= zero, one <= low


 👍 115 👎 0

*/

/*
题型: dp
题目: 统计构造好字符串的方案数
*/

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000007
	ans := 0

	g := gcd(zero, one)
	low = (low-1)/g + 1
	high /= g
	zero /= g
	one /= g

	mem := make([]int, high+1)
	mem[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			mem[i] = mem[i-zero]
		}
		if i >= one {
			mem[i] = (mem[i] + mem[i-one]) % mod
		}
		if i >= low {
			ans = (ans + mem[i]) % mod
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程: f(i) = f(i-zero) + f(i-one)
        low <= i <= high
边界状态: f(0) = 1; f(x) = 0, x < 0
*/

/*
递归:
func countGoodStrings(low int, high int, zero int, one int) int {
    var dfs func(int) int
    ans := 0
    dfs = func(i int) int {
        if i == 0 {
            return 1
        }
        if i<0 {
            return 0
        }
        return dfs(i-zero) + dfs(i-one)
    }
    for i := low;i<=high;i++ {
        ans += dfs(i)
    }
    return ans
}
时间复杂度: o(n^3)
空间复杂度: o(1)
*/

/*
记忆化搜索:
func countGoodStrings(low int, high int, zero int, one int) int {
    var dfs func(int) int
    mod := 1000000007
    ans := 0
    mem := make([]int,high+1)
    dfs = func(i int) int {
        if i == 0 {
            return 1
        }
        if i<0 {
            return 0
        }
        if mem[i] != 0 {
            return mem[i]
        }
        mem[i] = (dfs(i-zero) + dfs(i-one)) % mod
        return mem[i]
    }

    for i := low;i<=high;i++ {
        ans += dfs(i)
    }
    return ans % mod
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func countGoodStrings(low int, high int, zero int, one int) int {
    mod := 1000000007
    ans := 0
    mem := make([]int,high+1)
    mem[0] = 1
    for i := 1; i<=high;i++ {
        if i >= zero {
            mem[i] = mem[i-zero]
        }
        if i >= one {
            mem[i] += mem[i-one]
        }
        if i>= low {
            ans = (ans+mem[i]) % mod
        }
    }
    return ans
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
优化:
最大公约数优化：设zero和one的最大公约数为g，那么问题规模可以缩小为: low' = upflool(low/g), high' = downflool(high/g),
zero' = downflool(zero/g), one' = downflool(one/g)
func countGoodStrings(low int, high int, zero int, one int) int {
    mod := 1000000007
    ans := 0

    g := gcd(zero,one)
    low = (low-1)/g+1
    high /= g
    zero /= g
    one /= g

    mem := make([]int,high+1)
    mem[0] = 1
    for i := 1; i<=high;i++ {
        if i >= zero {
            mem[i] = mem[i-zero]
        }
        if i >= one {
            mem[i] = (mem[i] + mem[i-one])%mod
        }
        if i>= low {
            ans = (ans + mem[i])%mod
        }
    }
    return ans
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}
时间复杂度: o(high/g)
空间复杂度: o(high/g)
*/

func main() {

}
