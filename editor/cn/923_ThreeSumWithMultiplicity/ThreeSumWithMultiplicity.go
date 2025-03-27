package main

import "sort"

/**
给定一个整数数组
 arr ，以及一个整数 target 作为目标值，返回满足 i < j < k 且
 arr[i] + arr[j] + arr[k] == target 的元组 i, j, k 的数量。

 由于结果会非常大，请返回 10⁹ + 7 的模。



 示例 1：


输入：arr = [1,1,2,2,3,3,4,4,5,5], target = 8
输出：20
解释：
按值枚举(arr[i], arr[j], arr[k])：
(1, 2, 5) 出现 8 次；
(1, 3, 4) 出现 8 次；
(2, 2, 4) 出现 2 次；
(2, 3, 3) 出现 2 次。


 示例 2：


输入：arr = [1,1,2,2,2,2], target = 5
输出：12
解释：
arr[i] = 1, arr[j] = arr[k] = 2 出现 12 次：
我们从 [1,1] 中选择一个 1，有 2 种情况，
从 [2,2,2,2] 中选出两个 2，有 6 种情况。




 提示：


 3 <= arr.length <= 3000
 0 <= arr[i] <= 100
 0 <= target <= 300


 Related Topics 数组 哈希表 双指针 计数 排序 👍 152 👎 0

*/

/*
题型: 相向双指针
题目: 三数之和的多种可能
*/

// leetcode submit region begin(Prohibit modification and deletion)
// 双指针
func threeSumMulti(arr []int, target int) int {
	sort.Ints(arr) //排序

	res := 0
	//i<j<k
	//a[j]+a[k] = traget-a[i]
	for i := 0; i < len(arr)-2; i++ {
		j := i + 1
		k := len(arr) - 1
		for j < k {
			err := target - arr[i]
			if arr[j]+arr[k] == err {
				//fmt.Println(arr[i],arr[j],arr[k])
				// res++
				if arr[j] == arr[k] { //如果两个数相等
					res += (k - j + 1) * (k - j) / 2
					break
				} else {
					left := 1
					for arr[j] == arr[j+1] && j < k {
						//fmt.Println(arr[i],arr[j],arr[k])
						j++
						left++
					}
					right := 1
					for arr[k] == arr[k-1] && j < k {
						//fmt.Println(arr[i],arr[j],arr[k])
						k--
						right++
					}
					//fmt.Println(left,right)
					res += left * right
				}
				k--
			} else if arr[j]+arr[k] > err {
				k--
			} else {
				j++
			}

		}
	}
	return res % 1000000007

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
