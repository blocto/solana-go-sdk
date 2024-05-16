import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},t=s(`<h1 id="get-token-balance" tabindex="-1"><a class="header-anchor" href="#get-token-balance" aria-hidden="true">#</a> Get Token Balance</h1><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>

	<span class="token comment">// should pass a token account address</span>
	balance<span class="token punctuation">,</span> decimals<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetTokenAccountBalance</span><span class="token punctuation">(</span>
		context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token string">&quot;HeCBh32JJ8DxcjTyc6q46tirHR8hd2xj3mGoAcQ7eduL&quot;</span><span class="token punctuation">,</span>
	<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span><span class="token string">&quot;get balance error&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">// the smallest unit like lamports</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;balance&quot;</span><span class="token punctuation">,</span> balance<span class="token punctuation">)</span>
	<span class="token comment">// the decimals of mint which token account holds</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;decimals&quot;</span><span class="token punctuation">,</span> decimals<span class="token punctuation">)</span>

	<span class="token comment">// if you want use a normal unit, you can do</span>
	<span class="token comment">// balance / 10^decimals</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br></div></div>`,2);function p(e,c){return t}var l=n(a,[["render",p]]);export{l as default};
