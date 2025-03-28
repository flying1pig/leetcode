<p>给你一个字符串 <code>initialCurrency</code>，表示初始货币类型，并且你一开始拥有 <code>1.0</code> 单位的 <code>initialCurrency</code>。</p>

<p>另给你四个数组，分别表示货币对（字符串）和汇率（实数）：</p>

<ul> 
 <li><code>pairs1[i] = [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code> 表示在&nbsp;<strong>第 1 天</strong>，可以按照汇率 <code>rates1[i]</code> 将 <code>startCurrency<sub>i</sub></code> 转换为 <code>targetCurrency<sub>i</sub></code>。</li> 
 <li><code>pairs2[i] = [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code> 表示在&nbsp;<strong>第 2 天</strong>，可以按照汇率 <code>rates2[i]</code> 将 <code>startCurrency<sub>i</sub></code> 转换为 <code>targetCurrency<sub>i</sub></code>。</li> 
 <li>此外，每种 <code>targetCurrency</code> 都可以以汇率 <code>1 / rate</code> 转换回对应的 <code>startCurrency</code>。</li> 
</ul>

<p>你可以在&nbsp;<strong>第 1 天&nbsp;</strong>使用 <code>rates1</code> 进行任意次数的兑换（包括 0 次），然后在&nbsp;<strong>第 2 天&nbsp;</strong>使用 <code>rates2</code> 再进行任意次数的兑换（包括 0 次）。</p>

<p>返回在两天兑换后，最大可能拥有的 <code>initialCurrency</code> 的数量。</p>

<p><strong>注意：</strong>汇率是有效的，并且第 1 天和第 2 天的汇率之间相互独立，不会产生矛盾。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block"> 
 <p><strong>输入：</strong> <span class="example-io">initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.0,6.0]</span></p> 
</div>

<p><strong>输出：</strong> <span class="example-io">720.00000</span></p>

<p><strong>解释：</strong></p>

<p>根据题目要求，需要最大化最终的 <strong>EUR</strong> 数量，从 1.0 <strong>EUR</strong> 开始：</p>

<ul> 
 <li><strong>第 1 天：</strong> </li>
</ul>

    <ul>
    	<li>将 <strong>EUR</strong> 换成 <strong>USD</strong>，得到 2.0&nbsp;<strong>USD</strong>。</li>
    	<li>将 <strong>USD</strong> 换成 <strong>JPY</strong>，得到 6.0 <strong>JPY</strong>。</li>
    </ul>
    </li>
    <li><strong>第 2 天：</strong>
    <ul>
    	<li>将 <strong>JPY</strong> 换成 <strong>USD</strong>，得到 24.0 <strong>USD</strong>。</li>
    	<li>将 <strong>USD</strong> 换成 <strong>CHF</strong>，得到 120.0 <strong>CHF</strong>。</li>
    	<li>最后将 <strong>CHF</strong> 换回 <strong>EUR</strong>，得到 720.0 <strong>EUR</strong>。</li>
    </ul>
    </li>


<p><strong class="example">示例 2：</strong></p>

<div class="example-block"> 
 <p><strong>输入：</strong> <span class="example-io">initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 = [["NGN","EUR"]], rates2 = [6.0]</span></p> 
</div>

<p><strong>输出：</strong> <span class="example-io">1.50000</span></p>

<p><strong>解释：</strong></p>

<p>在第 1 天将 <strong>NGN</strong> 换成 <strong>EUR</strong>，并在第 2 天用反向汇率将 <strong>EUR</strong> 换回 <strong>NGN</strong>，可以最大化最终的 <strong>NGN</strong> 数量。</p>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block"> 
 <p><strong>输入：</strong> <span class="example-io">initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 = [["EUR","JPY"]], rates2 = [10.0]</span></p> 
</div>

<p><strong>输出：</strong> <span class="example-io">1.00000</span></p>

<p><strong>解释：</strong></p>

<p>在这个例子中，不需要在任何一天进行任何兑换。</p>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= initialCurrency.length &lt;= 3</code></li> 
 <li><code>initialCurrency</code> 仅由大写英文字母组成。</li> 
 <li><code>1 &lt;= n == pairs1.length &lt;= 10</code></li> 
 <li><code>1 &lt;= m == pairs2.length &lt;= 10</code></li> 
 <li><code>pairs1[i] == [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code></li> 
 <li><code>pairs2[i] == [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code></li> 
 <li><code>1 &lt;= startCurrency<sub>i</sub>.length, targetCurrency<sub>i</sub>.length &lt;= 3</code></li> 
 <li><code>startCurrency<sub>i</sub></code> 和 <code>targetCurrency<sub>i</sub></code> 仅由大写英文字母组成。</li> 
 <li><code>rates1.length == n</code></li> 
 <li><code>rates2.length == m</code></li> 
 <li><code>1.0 &lt;= rates1[i], rates2[i] &lt;= 10.0</code></li> 
 <li>输入保证两个转换图在各自的天数中没有矛盾或循环。</li> 
 <li>输入保证输出&nbsp;<strong>最大</strong>&nbsp;为&nbsp;<code>5 * 10<sup>10</sup></code>。</li> 
</ul>

<div><div>Related Topics</div><div><li>深度优先搜索</li><li>广度优先搜索</li><li>图</li><li>数组</li><li>字符串</li></div></div><br><div><li>👍 7</li><li>👎 0</li></div>