package main

/**
给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的
开始和结束，并且 intervals 按照 starti 升序排列。同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始
和结束。

 在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以
合并区间）。

 返回插入之后的 intervals。

 注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。



 示例 1：


输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]


 示例 2：


输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。




 提示：


 0 <= intervals.length <= 10⁴
 intervals[i].length == 2
 0 <= starti <= endi <= 10⁵
 intervals 根据 starti 按 升序 排列
 newInterval.length == 2
 0 <= start <= end <= 10⁵


 Related Topics 数组 👍 975 👎 0

*/

/*
题型: 贪心
题目: 插入区间
*/

// leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	i := 0
	intervalsLen := len(intervals)

	for i < intervalsLen && intervals[i][1] < newInterval[0] {
		i++
	}

	tempI := i
	for i < intervalsLen && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	intervals = append(intervals[:tempI], append([][]int{newInterval}, intervals[i:]...)...)
	return intervals
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
