import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},t=s(`<h1 id="get-mint" tabindex="-1"><a class="header-anchor" href="#get-mint" aria-hidden="true">#</a> Get Mint</h1><p>get detail from a exist mint</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/common&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/program/token&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">var</span> mintPubkey <span class="token operator">=</span> common<span class="token punctuation">.</span><span class="token function">PublicKeyFromString</span><span class="token punctuation">(</span><span class="token string">&quot;F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo&quot;</span><span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>

	getAccountInfoResponse<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetAccountInfo</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">TODO</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> mintPubkey<span class="token punctuation">.</span><span class="token function">ToBase58</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to get account info, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	mintAccount<span class="token punctuation">,</span> err <span class="token operator">:=</span> token<span class="token punctuation">.</span><span class="token function">MintAccountFromData</span><span class="token punctuation">(</span>getAccountInfoResponse<span class="token punctuation">.</span>Data<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to parse data to a mint account, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">&quot;%+v\\n&quot;</span><span class="token punctuation">,</span> mintAccount<span class="token punctuation">)</span>
	<span class="token comment">// {MintAuthority:9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde Supply:0 Decimals:8 IsInitialized:true FreezeAuthority:&lt;nil&gt;}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br></div></div>`,3);function p(o,e){return t}var l=n(a,[["render",p]]);export{l as default};
