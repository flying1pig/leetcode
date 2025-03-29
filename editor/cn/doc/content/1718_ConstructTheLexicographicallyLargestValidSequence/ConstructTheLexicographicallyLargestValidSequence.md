<p>给你一个整数&nbsp;<code>n</code>&nbsp;，请你找到满足下面条件的一个序列：</p>

<ul> 
 <li>整数&nbsp;<code>1</code>&nbsp;在序列中只出现一次。</li> 
 <li><code>2</code>&nbsp;到&nbsp;<code>n</code>&nbsp;之间每个整数都恰好出现两次。</li> 
 <li>对于每个&nbsp;<code>2</code>&nbsp;到&nbsp;<code>n</code>&nbsp;之间的整数&nbsp;<code>i</code>&nbsp;，两个&nbsp;<code>i</code>&nbsp;之间出现的距离恰好为&nbsp;<code>i</code>&nbsp;。</li> 
</ul>

<p>序列里面两个数 <code>a[i]</code>&nbsp;和 <code>a[j]</code>&nbsp;之间的 <strong>距离</strong>&nbsp;，我们定义为它们下标绝对值之差&nbsp;<code>|j - i|</code>&nbsp;。</p>

<p>请你返回满足上述条件中&nbsp;<strong>字典序最大</strong>&nbsp;的序列。题目保证在给定限制条件下，一定存在解。</p>

<p>一个序列&nbsp;<code>a</code>&nbsp;被认为比序列&nbsp;<code>b</code>&nbsp;（两者长度相同）字典序更大的条件是：&nbsp;<code>a</code> 和&nbsp;<code>b</code>&nbsp;中第一个不一样的数字处，<code>a</code>&nbsp;序列的数字比&nbsp;<code>b</code>&nbsp;序列的数字大。比方说，<code>[0,1,9,0]</code>&nbsp;比&nbsp;<code>[0,1,5,6]</code>&nbsp;字典序更大，因为第一个不同的位置是第三个数字，且&nbsp;<code>9</code>&nbsp;比&nbsp;<code>5</code>&nbsp;大。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>n = 3
<b>输出：</b>[3,1,2,3,2]
<b>解释：</b>[2,3,2,1,3] 也是一个可行的序列，但是 [3,1,2,3,2] 是字典序最大的序列。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>n = 5
<b>输出：</b>[5,3,1,4,3,5,2,4,2]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= n &lt;= 20</code></li> 
</ul>

<div><div>Related Topics</div><div><li>数组</li><li>回溯</li></div></div><br><div><li>👍 47</li><li>👎 0</li></div>