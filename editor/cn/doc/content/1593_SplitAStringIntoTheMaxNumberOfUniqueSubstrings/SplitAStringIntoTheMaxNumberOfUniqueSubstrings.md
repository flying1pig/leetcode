<p>给你一个字符串 <code>s</code> ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。</p>

<p>字符串 <code>s</code> 拆分后可以得到若干 <strong>非空子字符串</strong> ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 <strong>唯一的</strong> 。</p>

<p>注意：<strong>子字符串</strong> 是字符串中的一个连续字符序列。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>s = "ababccc"
<strong>输出：</strong>5
<strong>解释：</strong>一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = "aba"
<strong>输出：</strong>2
<strong>解释：</strong>一种最大拆分方法为 ['a', 'ba'] 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>s = "aa"
<strong>输出：</strong>1
<strong>解释：</strong>无法进一步拆分字符串。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li> <p><code>1 &lt;= s.length&nbsp;&lt;= 16</code></p> </li> 
 <li> <p><code>s</code> 仅包含小写英文字母</p> </li> 
</ul>

<div><div>Related Topics</div><div><li>哈希表</li><li>字符串</li><li>回溯</li></div></div><br><div><li>👍 65</li><li>👎 0</li></div>