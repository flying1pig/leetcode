package main

/**
一条街道上共有 n * 2 个 地块 ，街道的两侧各有 n 个地块。每一边的地块都按从 1 到 n 编号。每个地块上都可以放置一所房子。

 现要求街道同一侧不能存在两所房子相邻的情况，请你计算并返回放置房屋的方式数目。由于答案可能很大，需要对 10⁹ + 7 取余后再返回。

 注意，如果一所房子放置在这条街某一侧上的第 i 个地块，不影响在另一侧的第 i 个地块放置房子。



 示例 1：

 输入：n = 1
输出：4
解释：
可能的放置方式：
1. 所有地块都不放置房子。
2. 一所房子放在街道的某一侧。
3. 一所房子放在街道的另一侧。
4. 放置两所房子，街道两侧各放置一所。


 示例 2：
 输入：n = 2
输出：9
解释：如上图所示，共有 9 种可能的放置方式。




 提示：


 1 <= n <= 10⁴


 👍 69 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func countHousePlacements(n int) int {
	const mod = 1_000_000_007
	m0 := 1
	m1 := 2
	for i := 2; i <= n; i++ {
		newm := (m0 + m1) % mod
		m0 = m1
		m1 = newm
	}
	return m1 * m1 % mod
}

//leetcode submit region end(Prohibit modification and deletion)

/*
状态方程:
    第i个地块放置房子: f(i) = f(i-2)
    第i个地块不放置房子: f(i) = f(i-1)
    所以 f(i) = f(i-1) + f(i-2)
边界条件:
    f(0) = 1
    f(1) = 2
*/

/*
递归:
func countHousePlacements(n int) int {
    const mod = 1_000_000_007
    var f func(int) int
    f = func(i int) int {
        if i == 0 {
            return 1
        }
        if i==1 {
            return 2
        }
        return (f(i-1) + f(i-2))%mod
    }
    ans := f(n)
    return ans*ans%mod
}
时间复杂度: o(n^2)
空间复杂度: o(1)
*/

/*
记忆化搜索:
func countHousePlacements(n int) int {
    const mod = 1_000_000_007
    var f func(int) int
    mem := make([]int,n+1)
    f = func(i int) int {
        if i == 0 {
            return 1
        }
        if i==1 {
            return 2
        }
        if mem[i]==0 {
            mem[i] = (f(i-1) + f(i-2))%mod
        }
        return mem[i]
    }
    ans := f(n)
    return ans*ans%mod
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
递推:
func countHousePlacements(n int) int {
    const mod = 1_000_000_007
    mem := make([]int,n+1)
    mem[0] = 1
    mem[1] = 2
    for i:=2;i<=n;i++ {
        if mem[i] == 0 {
            mem[i] = (mem[i-1] + mem[i-2])%mod
        }
    }
    return mem[n]*mem[n]%mod
}
时间复杂度: o(n)
空间复杂度: o(n)
*/

/*
优化:

	func countHousePlacements(n int) int {
	    const mod = 1_000_000_007
	    m0 := 1
	    m1 := 2
	    for i:=2;i<=n;i++ {
	        newm := (m0+m1)%mod
	        m0 = m1
	        m1 = newm
	    }
	    return m1*m1%mod
	}

时间复杂度: o(n)
空间复杂度: o(1)
*/

func main() {

}
