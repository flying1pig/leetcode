package main

import (
	"fmt"
	"math"
)

/**
给定一个三角形 triangle ，找出自顶向下的最小路径和。

 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位
于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。



 示例 1：


输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。


 示例 2：


输入：triangle = [[-10]]
输出：-10




 提示：


 1 <= triangle.length <= 200
 triangle[0].length == 1
 triangle[i].length == triangle[i - 1].length + 1
 -10⁴ <= triangle[i][j] <= 10⁴




 进阶：


 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？


 👍 1428 👎 0

*/

/*
题型: dp
题目: 三角形最小路径和
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	mem := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mem[i] = make([]int, i+1)
		for j := range mem[i] {
			mem[i][j] = math.MaxInt
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if i == m-1 {
				mem[i][j] = triangle[i][j]
				continue
			}
			mem[i][j] = min(mem[i+1][j], mem[i+1][j+1]) + triangle[i][j]
		}
	}
	return mem[0][0]
}

//leetcode submit region end(Prohibit modification and deletion)

//mem[0][0] = min(mem[1][0],mem[1][1]) + triangle[0][0]
//mem[m-1][m-1] = min(mem[m][m-1],mem[m][m])
/*
定义f(i,j)为从等腰三角形顶点到底边的最小路径和
状态方程: f(i,j) = min(f(i-1,j),f(i-1,j-1)) + triangle[i][j]
必须从三角形底部走到三角形顶部，如果从三角形顶部走到三角形底部，可能会因遍历不全而得出错误答案。
边界条件:
	j<=i,i>=0,j>=0
	f(0,0) = triangle[0][0]
*/

/*
递归:
func minimumTotal(triangle [][]int) int {
	var f func(i,j int) int
	ans := math.MaxInt
	m := len(triangle)
	n := len(triangle[m-1])
	f = func(i, j int) int {
		if i<0 || j <0 || j>i{
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return triangle[0][0]
		}
		return min(f(i-1,j),f(i-1,j-1))+triangle[i][j]
	}
	for i:=0;i<n;i++ {
		ans = min(ans, f(m-1,i))
	}
	return ans
}
每次循环是一颗二叉树，树的高度是m，二叉树的叶子节点数为n^2，所以时间复杂度是o(n*m^2)
时间复杂度: o(n*m^2)
空间复杂度: o(n^2)
*/

/*
递推:
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	mem := make([][]int,m+1)
	for i:=0;i<=m;i++ {
		mem[i] = make([]int,i+1)
		for j := range mem[i] {
			mem[i][j] = math.MaxInt
		}
	}
	for i:= m-1;i>=0;i-- {
		for j :=i;j>=0;j-- {
			if i==m-1 {
				mem[i][j] = triangle[i][j]
				continue
			}
			mem[i][j] = min(mem[i+1][j],mem[i+1][j+1]) + triangle[i][j]
		}
	}
	return mem[0][0]
}
时间复杂度: o(n^2)
空间复杂度: o(n^2)
*/

/*
以下递归写法更简洁:

	func minimumTotal(triangle [][]int) int {
	    n := len(triangle)
	    f := make([][]int, n)
	    for i := range f {
	        f[i] = make([]int, i+1)
	    }
	    f[n-1] = triangle[n-1]
	    for i := n - 2; i >= 0; i-- {
	        for j, x := range triangle[i] {
	            f[i][j] = min(f[i+1][j], f[i+1][j+1]) + x
	        }
	    }
	    return f[0][0]
	}
*/
func main() {
	arr := [][]int{
		[]int{-10},
	}
	fmt.Println(minimumTotal(arr))
}
