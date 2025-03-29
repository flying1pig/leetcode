package main

/**
给你一个整数 n ，请你找到满足下面条件的一个序列：


 整数 1 在序列中只出现一次。
 2 到 n 之间每个整数都恰好出现两次。
 对于每个 2 到 n 之间的整数 i ，两个 i 之间出现的距离恰好为 i 。


 序列里面两个数 a[i] 和 a[j] 之间的 距离 ，我们定义为它们下标绝对值之差 |j - i| 。

 请你返回满足上述条件中 字典序最大 的序列。题目保证在给定限制条件下，一定存在解。

 一个序列 a 被认为比序列 b （两者长度相同）字典序更大的条件是： a 和 b 中第一个不一样的数字处，a 序列的数字比 b 序列的数字大。比方说，[0,1
,9,0] 比 [0,1,5,6] 字典序更大，因为第一个不同的位置是第三个数字，且 9 比 5 大。



 示例 1：

 输入：n = 3
输出：[3,1,2,3,2]
解释：[2,3,2,1,3] 也是一个可行的序列，但是 [3,1,2,3,2] 是字典序最大的序列。


 示例 2：

 输入：n = 5
输出：[5,3,1,4,3,5,2,4,2]




 提示：


 1 <= n <= 20


 Related Topics 数组 回溯 👍 47 👎 0

*/

/*
题型: 回溯
题目: 构建字典序最大的可行序列
*/

// leetcode submit region begin(Prohibit modification and deletion)
func constructDistancedSequence(n int) []int {
	seq := make([]int, 2*n-1) //结果序列
	state := 0                //第i位为1表示数字i+1已经使用过
	var dfs func(deep int) bool
	dfs = func(deep int) bool {
		if deep == 2*n-1 {
			return true
		}
		if seq[deep] != 0 { //位置deep已经用之前的数字填充过，直接递归枚举下一个位置
			return dfs(deep + 1)
		}
		for i := n - 1; i >= 0; i-- {
			if state>>i&1 == 0 { //如果第i位没有被使用
				state |= 1 << i //第i位置1
				seq[deep] = i + 1
				if i != 0 {
					if deep+i+1 >= 2*n-1 || seq[deep+i+1] != 0 {
						seq[deep] = 0
						state ^= 1 << i //第i位归0
						continue
					}
					seq[deep+i+1] = i + 1
				}
				if dfs(deep + 1) {
					return true
				}
				//如果代码走到这里，说明deep位置不能填充数字i，需要把deep和deep+i+1位置的数字都置0
				seq[deep] = 0
				if i != 0 {
					seq[deep+i+1] = 0
				}
				state ^= 1 << i //第i位归0
			}
		}
		return false
	}
	dfs(0)
	return seq
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
