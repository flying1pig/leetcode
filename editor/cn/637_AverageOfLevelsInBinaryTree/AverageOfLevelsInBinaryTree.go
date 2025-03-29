package main

/**
给定一个非空二叉树的根节点
 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10⁻⁵ 以内的答案可以被接受。



 示例 1：




输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。


 示例 2:




输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]




 提示：





 树中节点数量在 [1, 10⁴] 范围内
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 524 👎 0

*/

/*
题型: 二叉树层序遍历
题目: 二叉树的层平均值
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) (averages []float64) {
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		sum := 0
		curLevel := nextLevel
		nextLevel = nil
		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(len(curLevel)))
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
dfs
type data struct{ sum, count int }

func averageOfLevels(root *TreeNode) []float64 {
    levelData := []data{}
    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        if level < len(levelData) {
            levelData[level].sum += node.Val
            levelData[level].count++
        } else {
            levelData = append(levelData, data{node.Val, 1})
        }
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    dfs(root, 0)

    averages := make([]float64, len(levelData))
    for i, d := range levelData {
        averages[i] = float64(d.sum) / float64(d.count)
    }
    return averages
}

*/

func main() {

}
