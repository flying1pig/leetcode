package main

import "math"

/**
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。

 当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。


 例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。


 实现 StockSpanner 类：


 StockSpanner() 初始化类对象。
 int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。




 示例：


输入：
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
输出：
[null, 1, 1, 1, 2, 1, 4, 6]

解释：
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // 返回 1
stockSpanner.next(80);  // 返回 1
stockSpanner.next(60);  // 返回 1
stockSpanner.next(70);  // 返回 2
stockSpanner.next(60);  // 返回 1
stockSpanner.next(75);  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
stockSpanner.next(85);  // 返回 6




 提示：


 1 <= price <= 10⁵
 最多调用 next 方法 10⁴ 次


 Related Topics 栈 设计 数据流 单调栈 👍 489 👎 0

*/

/*
题型: 单调栈
题目: 股票价格跨度
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	day   int
	price int
}

type StockSpanner struct {
	stack  []pair
	curDay int // 第一个 next 调用算作第 0 天
}

func Constructor() StockSpanner {
	return StockSpanner{[]pair{{-1, math.MaxInt}}, -1} // 这样无需判断栈为空的情况
}

func (s *StockSpanner) Next(price int) int {
	for price >= s.stack[len(s.stack)-1].price {
		s.stack = s.stack[:len(s.stack)-1] // 栈顶数据后面不会再用到了，因为 price 更大
	}
	s.curDay++
	s.stack = append(s.stack, pair{s.curDay, price})
	return s.curDay - s.stack[len(s.stack)-2].day
}

//时间复杂度: o(1)
//空间复杂度：O(min(q,U))。其中 q 为 next 的调用次数，U 为 price 的范围

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
