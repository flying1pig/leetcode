package main

import (
	"strings"
)

/**
你是一位系统管理员，手里有一份文件夹列表 folder，你的任务是要删除该列表中的所有 子文件夹，并以 任意顺序 返回剩下的文件夹。

 如果文件夹 folder[i] 位于另一个文件夹 folder[j] 下，那么 folder[i] 就是 folder[j] 的 子文件夹 。folder[
j] 的子文件夹必须以 folder[j] 开头，后跟一个 "/"。例如，"/a/b" 是 "/a" 的一个子文件夹，但 "/b" 不是 "/a/b/c" 的一个
子文件夹。

 文件夹的「路径」是由一个或多个按以下格式串联形成的字符串：'/' 后跟一个或者多个小写英文字母。


 例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路径，而空字符串和 "/" 不是。




 示例 1：


输入：folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
输出：["/a","/c/d","/c/f"]
解释："/a/b" 是 "/a" 的子文件夹，而 "/c/d/e" 是 "/c/d" 的子文件夹。


 示例 2：


输入：folder = ["/a","/a/b/c","/a/b/d"]
输出：["/a"]
解释：文件夹 "/a/b/c" 和 "/a/b/d" 都会被删除，因为它们都是 "/a" 的子文件夹。


 示例 3：


输入: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
输出: ["/a/b/c","/a/b/ca","/a/b/d"]



 提示：


 1 <= folder.length <= 4 * 10⁴
 2 <= folder[i].length <= 100
 folder[i] 只包含小写字母和 '/'
 folder[i] 总是以字符 '/' 起始
 folder 每个元素都是 唯一 的


 Related Topics 深度优先搜索 字典树 数组 字符串 👍 171 👎 0

*/

/*
题型: 前缀树
题目: 删除子文件夹
*/

// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	children map[string]*Trie
	fid      int
}

func newTrie() *Trie {
	return &Trie{map[string]*Trie{}, -1}
}

func (this *Trie) insert(fid int, f string) {
	node := this
	ps := strings.Split(f, "/")
	for _, p := range ps[1:] {
		if _, ok := node.children[p]; !ok {
			node.children[p] = newTrie()
		}
		node = node.children[p]
	}
	node.fid = fid
}

func (this *Trie) search() (ans []int) {
	var dfs func(*Trie)
	dfs = func(root *Trie) {
		if root.fid != -1 {
			ans = append(ans, root.fid)
			return
		}
		for _, child := range root.children {
			dfs(child)
		}
	}
	dfs(this)
	return
}

func removeSubfolders(folder []string) (ans []string) {
	trie := newTrie()
	for i, f := range folder {
		trie.insert(i, f)
	}
	for _, i := range trie.search() {
		ans = append(ans, folder[i])
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/*
直接排序
func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}
*/

/*
字典树

	type Trie struct {
		children map[string]*Trie
		fid      int
	}

	func newTrie() *Trie {
		return &Trie{map[string]*Trie{}, -1}
	}

	func (this *Trie) insert(fid int, f string) {
		node := this
		ps := strings.Split(f, "/")
		for _, p := range ps[1:] {
			if _, ok := node.children[p]; !ok {
				node.children[p] = newTrie()
			}
			node = node.children[p]
		}
		node.fid = fid
	}

	func (this *Trie) search() (ans []int) {
		var dfs func(*Trie)
		dfs = func(root *Trie) {
			if root.fid != -1 {
				ans = append(ans, root.fid)
				return
			}
			for _, child := range root.children {
				dfs(child)
			}
		}
		dfs(this)
		return
	}

	func removeSubfolders(folder []string) (ans []string) {
		trie := newTrie()
		for i, f := range folder {
			trie.insert(i, f)
		}
		for _, i := range trie.search() {
			ans = append(ans, folder[i])
		}
		return
	}
*/
func main() {

}
