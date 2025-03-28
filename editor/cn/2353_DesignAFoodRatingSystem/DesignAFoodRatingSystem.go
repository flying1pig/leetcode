package main

import "container/heap"

/**
设计一个支持下述操作的食物评分系统：


 修改 系统中列出的某种食物的评分。
 返回系统中某一类烹饪方式下评分最高的食物。


 实现 FoodRatings 类：


 FoodRatings(String[] foods, String[] cuisines, int[] ratings) 初始化系统。食物由 foods、
cuisines 和 ratings 描述，长度均为 n 。



 foods[i] 是第 i 种食物的名字。
 cuisines[i] 是第 i 种食物的烹饪方式。
 ratings[i] 是第 i 种食物的最初评分。


 void changeRating(String food, int newRating) 修改名字为 food 的食物的评分。
 String highestRated(String cuisine) 返回指定烹饪方式 cuisine 下评分最高的食物的名字。如果存在并列，返回 字典序较
小 的名字。


 注意，字符串 x 的字典序比字符串 y 更小的前提是：x 在字典中出现的位置在 y 之前，也就是说，要么 x 是 y 的前缀，或者在满足 x[i] != y[
i] 的第一个位置 i 处，x[i] 在字母表中出现的位置在 y[i] 之前。



 示例：

 输入
["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated",
"changeRating", "highestRated"]
[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean",
"japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], [
"korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]
输出
[null, "kimchi", "ramen", null, "sushi", null, "ramen"]

解释
FoodRatings foodRatings = new FoodRatings(["kimchi", "miso", "sushi",
"moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese",
"korean"], [9, 12, 8, 15, 14, 7]);
foodRatings.highestRated("korean"); // 返回 "kimchi"
                                    // "kimchi" 是分数最高的韩式料理，评分为 9 。
foodRatings.highestRated("japanese"); // 返回 "ramen"
                                      // "ramen" 是分数最高的日式料理，评分为 14 。
foodRatings.changeRating("sushi", 16); // "sushi" 现在评分变更为 16 。
foodRatings.highestRated("japanese"); // 返回 "sushi"
                                      // "sushi" 是分数最高的日式料理，评分为 16 。
foodRatings.changeRating("ramen", 16); // "ramen" 现在评分变更为 16 。
foodRatings.highestRated("japanese"); // 返回 "ramen"
                                      // "sushi" 和 "ramen" 的评分都是 16 。
                                      // 但是，"ramen" 的字典序比 "sushi" 更小。




 提示：


 1 <= n <= 2 * 10⁴
 n == foods.length == cuisines.length == ratings.length
 1 <= foods[i].length, cuisines[i].length <= 10
 foods[i]、cuisines[i] 由小写英文字母组成
 1 <= ratings[i] <= 10⁸
 foods 中的所有字符串 互不相同
 在对 changeRating 的所有调用中，food 是系统中食物的名字。
 在对 highestRated 的所有调用中，cuisine 是系统中 至少一种 食物的烹饪方式。
 最多调用 changeRating 和 highestRated 总计 2 * 10⁴ 次


 Related Topics 设计 数组 哈希表 字符串 有序集合 堆（优先队列） 👍 63 👎 0

*/

/*
题型: 堆
题目: 设计食物评分系统
*/

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	rating int
	s      string
}

type FoodRatings struct {
	foodMap    map[string]pair
	cuisineMap map[string]*hp
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	foodMap := map[string]pair{}
	cuisineMap := map[string]*hp{}
	for i, food := range foods {
		rating, cuisine := ratings[i], cuisines[i]
		foodMap[food] = pair{rating, cuisine}
		if cuisineMap[cuisine] == nil {
			cuisineMap[cuisine] = &hp{}
		}
		heap.Push(cuisineMap[cuisine], pair{rating, food})
	}
	return FoodRatings{foodMap, cuisineMap}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.foodMap[food]
	// 直接添加新数据，后面 HighestRated 再删除旧的
	heap.Push(r.cuisineMap[p.s], pair{newRating, food})
	p.rating = newRating
	r.foodMap[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	h := r.cuisineMap[cuisine]
	// 堆顶的食物评分不等于其实际值
	for h.Len() > 0 && (*h)[0].rating != r.foodMap[(*h)[0].s].rating {
		heap.Pop(h)
	}
	return (*h)[0].s
}

type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.rating > b.rating || a.rating == b.rating && a.s < b.s
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
