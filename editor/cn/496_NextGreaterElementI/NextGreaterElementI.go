package main

/**
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2
[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。



 示例 1：


输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

 示例 2：


输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。




 提示：


 1 <= nums1.length <= nums2.length <= 1000
 0 <= nums1[i], nums2[i] <= 10⁴
 nums1和nums2中所有整数 互不相同
 nums1 中的所有整数同样出现在 nums2 中




 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？

 Related Topics 栈 数组 哈希表 单调栈 👍 1245 👎 0

*/

/*
题型: 单调栈
题目: 下一个更大元素 I
*/

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1, nums2 []int) []int {
	idx := make(map[int]int, len(nums1)) // 预分配空间
	for i, x := range nums1 {
		idx[x] = i
	}
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for _, x := range nums2 {
		for len(st) > 0 && x > st[len(st)-1] {
			// x 是栈顶的下一个更大元素
			// 既然栈顶已经算出答案，弹出
			ans[idx[st[len(st)-1]]] = x // 记录答案
			st = st[:len(st)-1]
		}
		if _, ok := idx[x]; ok { // x 在 nums1 中
			st = append(st, x) // 只需把在 nums1 中的元素入栈
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
从右到左
func nextGreaterElement(nums1, nums2 []int) []int {
    idx := make(map[int]int, len(nums1)) // 预分配空间
    for i, x := range nums1 {
        idx[x] = i
    }
    ans := make([]int, len(nums1))
    for i := range ans {
        ans[i] = -1
    }
    st := []int{}
    for i := len(nums2) - 1; i >= 0; i-- {
        x := nums2[i]
        for len(st) > 0 && x >= st[len(st)-1] {
            // 由于 x 的出现，栈顶元素永远不会是左边元素的「下一个更大元素」
            st = st[:len(st)-1]
        }
        if len(st) > 0 { // x 在 nums1 中
            if j, ok := idx[x]; ok {
                ans[j] = st[len(st)-1] // 记录答案
            }
        }
        st = append(st, x)
    }
    return ans
}
时间复杂度: 时间复杂度：O(n+m)，其中 n 是nums1的长度，m 是nums2的长度
空间复杂度: o(n+m)
*/

/*
从左到右
func nextGreaterElement(nums1, nums2 []int) []int {
    idx := make(map[int]int, len(nums1)) // 预分配空间
    for i, x := range nums1 {
        idx[x] = i
    }
    ans := make([]int, len(nums1))
    for i := range ans {
        ans[i] = -1
    }
    st := []int{}
    for _, x := range nums2 {
        for len(st) > 0 && x > st[len(st)-1] {
            // x 是栈顶的下一个更大元素
            // 既然栈顶已经算出答案，弹出
            ans[idx[st[len(st)-1]]] = x // 记录答案
            st = st[:len(st)-1]
        }
        if _, ok := idx[x]; ok { // x 在 nums1 中
            st = append(st, x) // 只需把在 nums1 中的元素入栈
        }
    }
    return ans
}
时间复杂度：O(n+m)
空间复杂度：O(n)
*/

func main() {

}
