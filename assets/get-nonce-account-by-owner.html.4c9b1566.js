import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},t=s(`<h1 id="get-nonce-account-by-owner" tabindex="-1"><a class="header-anchor" href="#get-nonce-account-by-owner" aria-hidden="true">#</a> Get Nonce Account By Owner</h1><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;encoding/base64&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/common&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/program/system&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>

	res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span>RpcClient<span class="token punctuation">.</span><span class="token function">GetProgramAccountsWithConfig</span><span class="token punctuation">(</span>
		context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		common<span class="token punctuation">.</span>SystemProgramID<span class="token punctuation">.</span><span class="token function">ToBase58</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		rpc<span class="token punctuation">.</span>GetProgramAccountsConfig<span class="token punctuation">{</span>
			Encoding<span class="token punctuation">:</span> rpc<span class="token punctuation">.</span>AccountEncodingBase64<span class="token punctuation">,</span>
			Filters<span class="token punctuation">:</span> <span class="token punctuation">[</span><span class="token punctuation">]</span>rpc<span class="token punctuation">.</span>GetProgramAccountsConfigFilter<span class="token punctuation">{</span>
				<span class="token punctuation">{</span>
					DataSize<span class="token punctuation">:</span> system<span class="token punctuation">.</span>NonceAccountSize<span class="token punctuation">,</span>
				<span class="token punctuation">}</span><span class="token punctuation">,</span>
				<span class="token punctuation">{</span>
					MemCmp<span class="token punctuation">:</span> <span class="token operator">&amp;</span>rpc<span class="token punctuation">.</span>GetProgramAccountsConfigFilterMemCmp<span class="token punctuation">{</span>
						Offset<span class="token punctuation">:</span> <span class="token number">8</span><span class="token punctuation">,</span>
						Bytes<span class="token punctuation">:</span>  <span class="token string">&quot;9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde&quot;</span><span class="token punctuation">,</span> <span class="token comment">// owner address</span>
					<span class="token punctuation">}</span><span class="token punctuation">,</span>
				<span class="token punctuation">}</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">,</span>
	<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to get program accounts, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> a <span class="token operator">:=</span> <span class="token keyword">range</span> res<span class="token punctuation">.</span>Result <span class="token punctuation">{</span>
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;pubkey&quot;</span><span class="token punctuation">,</span> a<span class="token punctuation">.</span>Pubkey<span class="token punctuation">)</span>
		data<span class="token punctuation">,</span> err <span class="token operator">:=</span> base64<span class="token punctuation">.</span>StdEncoding<span class="token punctuation">.</span><span class="token function">DecodeString</span><span class="token punctuation">(</span><span class="token punctuation">(</span>a<span class="token punctuation">.</span>Account<span class="token punctuation">.</span>Data<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span>any<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to decode data, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		nonceAccount<span class="token punctuation">,</span> err <span class="token operator">:=</span> system<span class="token punctuation">.</span><span class="token function">NonceAccountDeserialize</span><span class="token punctuation">(</span>data<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to parse nonce account, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">&quot;%+v\\n&quot;</span><span class="token punctuation">,</span> nonceAccount<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br><span class="line-number">32</span><br><span class="line-number">33</span><br><span class="line-number">34</span><br><span class="line-number">35</span><br><span class="line-number">36</span><br><span class="line-number">37</span><br><span class="line-number">38</span><br><span class="line-number">39</span><br><span class="line-number">40</span><br><span class="line-number">41</span><br><span class="line-number">42</span><br><span class="line-number">43</span><br><span class="line-number">44</span><br><span class="line-number">45</span><br><span class="line-number">46</span><br><span class="line-number">47</span><br><span class="line-number">48</span><br><span class="line-number">49</span><br><span class="line-number">50</span><br><span class="line-number">51</span><br><span class="line-number">52</span><br></div></div>`,2);function p(o,c){return t}var u=n(a,[["render",p]]);export{u as default};
