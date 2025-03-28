package main

import "container/heap"

/**
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近 10 条推文。

 实现 Twitter 类：


 Twitter() 初始化简易版推特对象
 void postTweet(int userId, int tweetId) 根据给定的 tweetId 和 userId 创建一条新推文。每次调用此函数都
会使用一个不同的 tweetId 。
 List<Integer> getNewsFeed(int userId) 检索当前用户新闻推送中最近 10 条推文的 ID 。新闻推送中的每一项都必须是由用
户关注的人或者是用户自己发布的推文。推文必须 按照时间顺序由最近到最远排序 。
 void follow(int followerId, int followeeId) ID 为 followerId 的用户开始关注 ID 为
followeeId 的用户。
 void unfollow(int followerId, int followeeId) ID 为 followerId 的用户不再关注 ID 为
followeeId 的用户。




 示例：


输入
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed",
"unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
输出
[null, null, [5], null, null, [6, 5], null, [5]]

解释
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // 用户 1 发送了一条新推文 (用户 id = 1, 推文 id = 5)
twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含一个 id 为 5 的推文
twitter.follow(1, 2);    // 用户 1 关注了用户 2
twitter.postTweet(2, 6); // 用户 2 发送了一个新推文 (推文 id = 6)
twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含两个推文，id 分别为 -> [6, 5] 。推文 id
6 应当在推文 id 5 之前，因为它是在 5 之后发送的
twitter.unfollow(1, 2);  // 用户 1 取消关注了用户 2
twitter.getNewsFeed(1);  // 用户 1 获取推文应当返回一个列表，其中包含一个 id 为 5 的推文。因为用户 1 已经不再关注用户
2



 提示：


 1 <= userId, followerId, followeeId <= 500
 0 <= tweetId <= 10⁴
 所有推特的 ID 都互不相同
 postTweet、getNewsFeed、follow 和 unfollow 方法最多调用 3 * 10⁴ 次
 用户不能关注自己


 Related Topics 设计 哈希表 链表 堆（优先队列） 👍 408 👎 0

*/

/*
题型: 堆
题目: 设计推特
*/

// leetcode submit region begin(Prohibit modification and deletion)
type Twitter struct {
	users map[int]*user // 全局用户列表，快速查到用户，没有就新建一个user 实例
	time  int           // 全局时间戳
}

type user struct {
	userId    int
	followee  map[int]*user
	tweetList *tweetText // 发表的文章列表，链表的形式（合并k 个有序链表）
}

type tweetText struct {
	tweetId int
	time    int
	next    *tweetText // 指针
}

type tweetTextHeap []*tweetText

func (h tweetTextHeap) Len() int           { return len(h) }
func (h tweetTextHeap) Less(i, j int) bool { return h[i].time > h[j].time } // 按照时间顺序，新发的推特id 越大越靠前
func (h tweetTextHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tweetTextHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*tweetText))
}

func (h *tweetTextHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Constructor /** Initialize your data structure here. */
func Constructor() Twitter {
	t := Twitter{
		users: make(map[int]*user),
	}
	return t
}

// PostTweet /** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	text := tweetText{
		tweetId: tweetId,
		time:    t.getTime(),
	}
	u := t.getUser(userId)
	text.next = u.tweetList // 第三个属性赋值，将next 指向当前用户的列表头结点
	u.tweetList = &text     // 头结点移动到本次text的位置，给下一次发推做准备 text.next -> head ,head -> text,尾插法
}

// GetNewsFeed /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	u := t.getUser(userId)
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	texts := make(tweetTextHeap, 0, 1+len(u.followee))
	if u.tweetList != nil {
		texts = append(texts, u.tweetList)
	}
	for _, v := range u.followee {
		if v.tweetList != nil {
			texts = append(texts, v.tweetList)
		}
	}
	// 所有用户的推特都加入了texts
	res := make([]int, 0, 10)

	h := &texts
	heap.Init(h)
	for len(res) < 10 && h.Len() > 0 {
		text := heap.Pop(h).(*tweetText)
		res = append(res, text.tweetId)
		if text.next != nil {
			heap.Push(h, text.next)
		}
	}
	return res
}

// Follow /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	u1 := t.getUser(followerId)
	u2 := t.getUser(followeeId)
	u1.followee[followeeId] = u2 //关注列表，当前用户关注了哪些用户
}

// Unfollow /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */ 取关的用户已经注销了，或者
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId { //自己不能取关自己
		return
	}
	u1 := t.getUser(followerId)
	_ = t.getUser(followeeId)
	delete(u1.followee, followeeId)
}

func (t *Twitter) getUser(userId int) *user {
	u, ok := t.users[userId]
	if !ok {
		u = &user{
			userId:   userId,
			followee: make(map[int]*user),
		}
		t.users[userId] = u
	}
	return u
}

func (t *Twitter) getTime() int {
	t.time++
	return t.time
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
