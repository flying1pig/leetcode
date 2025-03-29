package main

/**
请你设计一个迭代器类 CombinationIterator ，包括以下内容：


 CombinationIterator(string characters, int combinationLength) 一个构造函数，输入参数包括：用一个
 有序且字符唯一 的字符串 characters（该字符串只包含小写英文字母）和一个数字 combinationLength 。
 函数 next() ，按 字典序 返回长度为 combinationLength 的下一个字母组合。
 函数 hasNext() ，只有存在长度为 combinationLength 的下一个字母组合时，才返回 true




 示例 1：


输入:
["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"]

[["abc", 2], [], [], [], [], [], []]
输出：
[null, "ab", true, "ac", true, "bc", false]
解释：
CombinationIterator iterator = new CombinationIterator("abc", 2); // 创建迭代器
iterator
iterator.next(); // 返回 "ab"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "ac"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "bc"
iterator.hasNext(); // 返回 false




 提示：


 1 <= combinationLength <= characters.length <= 15
 characters 中每个字符都 不同
 每组测试数据最多对 next 和 hasNext 调用 10⁴次
 题目保证每次调用函数 next 时都存在下一个字母组合。


 Related Topics 设计 字符串 回溯 迭代器 👍 78 👎 0

*/

/*
题型: 回溯
题目: 字母组合迭代器
*/

// leetcode submit region begin(Prohibit modification and deletion)
type CombinationIterator struct {
	Res  []string
	Nums int //Res的数量
	Now  int //现在的位置
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var all CombinationIterator
	all.Res = make([]string, 0)
	a := make([]byte, 0)
	dfs(a, characters, 0, combinationLength, &all)
	all.Nums = len(all.Res)
	all.Now = 0
	return all
}
func dfs(chance []byte, characters string, num int, combinationLength int, res *CombinationIterator) {
	if len(chance) == combinationLength { //当长度符合要求后  将答案存入Res 返回
		ll := make([]byte, len(chance))
		copy(ll, chance)
		res.Res = append(res.Res, string(ll))
		res.Nums += 1
		return
	}
	for i := 0; i < len(characters); i++ {
		if num != 0 {
			if chance[num-1] >= characters[i] { //如果当前字符小于chance中的前一个字符 剪枝（字典序）
				continue
			}
		}
		chance = append(chance, characters[i])
		dfs(chance, characters, num+1, combinationLength, res)
		chance = chance[:len(chance)-1]
	}
	return
}
func (this *CombinationIterator) Next() string {
	res := this.Res[this.Now]
	this.Now += 1
	return res

}

func (this *CombinationIterator) HasNext() bool {
	if this.Now < this.Nums {
		return true
	}
	return false
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
