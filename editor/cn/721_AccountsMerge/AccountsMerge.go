package main

import "slices"

/**
给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元
素是 emails 表示该账户的邮箱地址。

 现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们
可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

 合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。



 示例 1：


输入：accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John",
"johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], [
"Mary", "mary@mail.com"]]
输出：[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],
 ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
解释：
第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。
第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。


 示例 2：


输入：accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@
m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@
m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co",
"Fern1@m.co","Fern0@m.co"]]
输出：[["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co",
"Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],[
"Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.
co","Fern5@m.co"]]




 提示：


 1 <= accounts.length <= 1000
 2 <= accounts[i].length <= 10
 1 <= accounts[i][j].length <= 30
 accounts[i][0] 由英文字母组成
 accounts[i][j] (for j > 0) 是有效的邮箱地址


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 哈希表 字符串 排序 👍 562 👎 0

*/

/*
题型: 图论dfs
题目: 账户合并
*/

// leetcode submit region begin(Prohibit modification and deletion)
func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIdx := map[string][]int{}
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToIdx[email] = append(emailToIdx[email], i)
		}
	}

	vis := make([]bool, len(accounts))
	emailSet := map[string]struct{}{} // 用于收集 DFS 中访问到的邮箱地址
	var dfs func(int)
	dfs = func(i int) {
		vis[i] = true
		for _, email := range accounts[i][1:] { // 遍历 i 的所有邮箱地址
			if _, has := emailSet[email]; has {
				continue
			}
			emailSet[email] = struct{}{}
			for _, j := range emailToIdx[email] { // 遍历所有包含该邮箱地址的账户下标 j
				if !vis[j] { // j 没有访问过
					dfs(j)
				}
			}
		}
	}

	for i, b := range vis {
		if b {
			continue
		}
		clear(emailSet)
		dfs(i)

		res := make([]string, 1, len(emailSet)+1)
		res[0] = accounts[i][0]
		for email := range emailSet {
			res = append(res, email)
		}
		slices.Sort(res[1:])

		ans = append(ans, res)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/*
思路:
1. 把 accounts 中的信息提取到哈希表 emailToIdx 中，key 为邮箱地址，value 为这个邮箱对应的账户下标列表。
2. 初始化一个长为 n 的全为 false 的布尔数组 vis，用来标记访问过的账户下标。
3. 遍历 vis，如果 i 没有访问过，即 vis[i]=false，则从 i 开始 DFS。
	a. DFS 之前，创建一个哈希集合 emails，用来保存 DFS 中访问到的邮箱地址。
	b. 开始 DFS。首先标记 vis[i]=true。
	c. 遍历 accounts[i] 的邮箱地址 email。
	d. 如果 email 在哈希集合 emails 中，则跳过；否则把 email 加入哈希集合 emails。
	e. 遍历 emailToIdx[email]，也就是所有包含该邮箱地址的账户下标 j，如果 j 没有访问过，即 vis[j]=false，则继续 DFS j。
4. DFS 结束后，把 emails 中的元素按照字典序从小到大排序，然后和 accounts[i][0] 一起加入答案。
5. 返回答案。

*/

func main() {

}
