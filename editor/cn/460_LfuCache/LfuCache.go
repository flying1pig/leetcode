package main

import "container/list"

/**
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

 实现 LFUCache 类：


 LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
 int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
 void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。


 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



 示例：


输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
"get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3



 提示：


 1 <= capacity <= 10⁴
 0 <= key <= 10⁵
 0 <= value <= 10⁹
 最多调用 2 * 10⁵ 次 get 和 put 方法


 Related Topics 设计 哈希表 链表 双向链表 👍 871 👎 0

*/

/*
题型: 链表
题目: LFU 缓存
*/

// leetcode submit region begin(Prohibit modification and deletion)
type entry struct {
	key, value, freq int // freq 表示这本书被看的次数
}

type LFUCache struct {
	capacity   int
	minFreq    int
	keyToNode  map[int]*list.Element
	freqToList map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		keyToNode:  map[int]*list.Element{},
		freqToList: map[int]*list.List{},
	}
}

func (c *LFUCache) pushFront(e *entry) {
	if _, ok := c.freqToList[e.freq]; !ok {
		c.freqToList[e.freq] = list.New() // 双向链表
	}
	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
}

func (c *LFUCache) getEntry(key int) *entry {
	node := c.keyToNode[key]
	if node == nil { // 没有这本书
		return nil
	}
	e := node.Value.(*entry)
	lst := c.freqToList[e.freq]
	lst.Remove(node)    // 把这本书抽出来
	if lst.Len() == 0 { // 抽出来后，这摞书是空的
		delete(c.freqToList, e.freq) // 移除空链表
		if c.minFreq == e.freq {     // 这摞书是最左边的
			c.minFreq++
		}
	}
	e.freq++       // 看书次数 +1
	c.pushFront(e) // 放在右边这摞书的最上面
	return e
}

func (c *LFUCache) Get(key int) int {
	if e := c.getEntry(key); e != nil { // 有这本书
		return e.value
	}
	return -1 // 没有这本书
}

func (c *LFUCache) Put(key, value int) {
	if e := c.getEntry(key); e != nil { // 有这本书
		e.value = value // 更新 value
		return
	}
	if len(c.keyToNode) == c.capacity { // 书太多了
		lst := c.freqToList[c.minFreq]                           // 最左边那摞书
		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key) // 移除这摞书的最下面的书
		if lst.Len() == 0 {                                      // 这摞书是空的
			delete(c.freqToList, c.minFreq) // 移除空链表
		}
	}
	c.pushFront(&entry{key, value, 1}) // 新书放在「看过 1 次」的最上面
	c.minFreq = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
