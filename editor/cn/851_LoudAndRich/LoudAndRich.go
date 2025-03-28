package main

/**
有一组 n 个人作为实验对象，从 0 到 n - 1 编号，其中每个人都有不同数目的钱，以及不同程度的安静值（quietness）。为了方便起见，我们将编号为
x 的人简称为 "person x "。

 给你一个数组 richer ，其中 richer[i] = [ai, bi] 表示 person ai 比 person bi 更有钱。另给你一个整数数组
quiet ，其中 quiet[i] 是 person i 的安静值。richer 中所给出的数据 逻辑自洽（也就是说，在 person x 比 person
y 更有钱的同时，不会出现 person y 比 person x 更有钱的情况 ）。

 现在，返回一个整数数组 answer 作为答案，其中 answer[x] = y 的前提是，在所有拥有的钱肯定不少于 person x 的人中，person
y 是最不安静的人（也就是安静值 quiet[y] 最小的人）。



 示例 1：


输入：richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,
0]
输出：[5,5,2,5,4,5,6,7]
解释：
answer[0] = 5，
person 5 比 person 3 有更多的钱，person 3 比 person 1 有更多的钱，person 1 比 person 0 有更多的钱。
唯一较为安静（有较低的安静值 quiet[x]）的人是 person 7，
但是目前还不清楚他是否比 person 0 更有钱。
answer[7] = 7，
在所有拥有的钱肯定不少于 person 7 的人中（这可能包括 person 3，4，5，6 以及 7），
最安静（有较低安静值 quiet[x]）的人是 person 7。
其他的答案也可以用类似的推理来解释。


 示例 2：


输入：richer = [], quiet = [0]
输出：[0]




 提示：


 n == quiet.length
 1 <= n <= 500
 0 <= quiet[i] < n
 quiet 的所有值 互不相同
 0 <= richer.length <= n * (n - 1) / 2
 0 <= ai, bi < n
 ai != bi
 richer 中的所有数对 互不相同
 对 richer 的观察在逻辑上是一致的


 Related Topics 深度优先搜索 图 拓扑排序 数组 👍 259 👎 0

*/

/*
题型: 图论拓扑排序
题目: 喧闹和富有
*/

// leetcode submit region begin(Prohibit modification and deletion)
func loudAndRich(richer [][]int, quiet []int) []int {
	grah := make([][]int, len(quiet))
	// 建图，有钱人指向没钱人,并记录没钱人的入度，有几个人比他钱多
	in := make([]int, len(quiet))
	for _, i := range richer {
		grah[i[0]] = append(grah[i[0]], i[1])
		in[i[1]]++
	}
	res := make([]int, len(quiet))
	// 最有钱的人的ans就是自己，初始化默认值
	for i := range res {
		res[i] = i
	}
	// 初始化队列
	q := make([]int, 0, len(quiet))
	// 先把最有钱的人放进队列，也就是入度为0的
	for p, i := range in {
		if i == 0 {
			q = append(q, p)
		}
	}
	// 遍历队列
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// 遍历比他钱少的
		for _, i := range grah[cur] {
			// 如果钱最多的安静值，比比他钱少的安静值少，就更新钱少的ans
			if quiet[res[cur]] < quiet[res[i]] {
				res[i] = res[cur]
			}
			// 钱少的入度减1，说明这个钱多的已经处理完
			in[i]--
			// 当入度为0的时候，说明这个就是钱最多的了，放入队列处理
			if in[i] == 0 {
				q = append(q, i)
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
