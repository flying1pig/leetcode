package main

import "sort"

/**
在一条单行道上，有 n 辆车开往同一目的地。目的地是几英里以外的 target 。

 给定两个整数数组 position 和 speed ，长度都是 n ，其中 position[i] 是第 i 辆车的位置， speed[i] 是第 i 辆车的
速度(单位是英里/小时)。

 一辆车永远不会超过前面的另一辆车，但它可以追上去，并以较慢车的速度在另一辆车旁边行驶。

 车队 是指并排行驶的一辆或几辆汽车。车队的速度是车队中 最慢 的车的速度。

 即便一辆车在 target 才赶上了一个车队，它们仍然会被视作是同一个车队。

 返回到达目的地的车队数量 。



 示例 1：


 输入：target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]


 输出：3

 解释：


 从 10（速度为 2）和 8（速度为 4）开始的车会组成一个车队，它们在 12 相遇。车队在 target 形成。
 从 0（速度为 1）开始的车不会追上其它任何车，所以它自己是一个车队。
 从 5（速度为 1） 和 3（速度为 3）开始的车组成一个车队，在 6 相遇。车队以速度 1 移动直到它到达 target。


 示例 2：


 输入：target = 10, position = [3], speed = [3]


 输出：1

 解释： 只有一辆车，因此只有一个车队。

 示例 3：


 输入：target = 100, position = [0,2,4], speed = [4,2,1]


 输出：1

 解释：


 从 0（速度为 4） 和 2（速度为 2）开始的车组成一个车队，在 4 相遇。从 4 开始的车（速度为 1）移动到了 5。
 然后，在 4（速度为 2）的车队和在 5（速度为 1）的车成为一个车队，在 6 相遇。车队以速度 1 移动直到它到达 target。




 提示：


 n == position.length == speed.length
 1 <= n <= 10⁵
 0 < target <= 10⁶
 0 <= position[i] < target
 position 中每个值都 不同
 0 < speed[i] <= 10⁶


 Related Topics 栈 数组 排序 单调栈 👍 220 👎 0

*/

/*
题型: 单调栈
题目: 车队
*/

// leetcode submit region begin(Prohibit modification and deletion)
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	// 创建一个结构体来存储位置和速度
	cars := make([]struct {
		pos int
		spd int
	}, n)

	for i := range position {
		cars[i] = struct {
			pos int
			spd int
		}{position[i], speed[i]}
	}

	// 按照位置从大到小排序
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	// 计算每辆车的到达时间
	times := make([]float64, n)
	for i := range cars {
		times[i] = float64(target-cars[i].pos) / float64(cars[i].spd)
	}

	// 使用栈来记录车队的到达时间
	stack := []float64{}
	for _, time := range times {
		// 如果当前车的到达时间大于栈顶的到达时间，说明形成新的车队
		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}

/*
单调栈进阶题目:
1019. 链表中的下一个更大节点 1571
962. 最大宽度坡 1608
1124. 表现良好的最长时间段 1908
456. 132 模式 约 2000
3113. 边界元素是最大值的子数组数目 2046
2866. 美丽塔 II 2072
1944. 队列中可以看到的人数 2105
2454. 下一个更大元素 IV 2175
1130. 叶值的最小代价生成树 O(n) 做法
2289. 使数组按非递减顺序排列 2482
1776. 车队 II 2531
3420. 统计 K 次操作以内得到非递减子数组的数目 2855 树形结构
3221. 最大数组跳跃得分 II（会员题）
1966. 未排序数组中的可被二分搜索的数（会员题）
2832. 每个元素为最大值的最大范围（会员题）
2282. 在一个网格中可以看到的人数（会员题）
*/
