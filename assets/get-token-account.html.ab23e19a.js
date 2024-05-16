import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},t=s(`<h1 id="get-token-account" tabindex="-1"><a class="header-anchor" href="#get-token-account" aria-hidden="true">#</a> Get Token Account</h1><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/program/token&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>

	<span class="token comment">// token account address</span>
	getAccountInfoResponse<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetAccountInfo</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">TODO</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token string">&quot;HeCBh32JJ8DxcjTyc6q46tirHR8hd2xj3mGoAcQ7eduL&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to get account info, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	tokenAccount<span class="token punctuation">,</span> err <span class="token operator">:=</span> token<span class="token punctuation">.</span><span class="token function">TokenAccountFromData</span><span class="token punctuation">(</span>getAccountInfoResponse<span class="token punctuation">.</span>Data<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to parse data to a token account, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">&quot;%+v\\n&quot;</span><span class="token punctuation">,</span> tokenAccount<span class="token punctuation">)</span>
	<span class="token comment">// {Mint:F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo Owner:9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde Amount:100000000 Delegate:&lt;nil&gt; State:1 IsNative:&lt;nil&gt; DelegatedAmount:0 CloseAuthority:&lt;nil&gt;}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br></div></div>`,2);function p(o,e){return t}var l=n(a,[["render",p]]);export{l as default};
