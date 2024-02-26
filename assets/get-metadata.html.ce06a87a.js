import{_ as n,e as a}from"./app.aa4fcc9f.js";const s={},t=a(`<h1 id="get-token-metadata" tabindex="-1"><a class="header-anchor" href="#get-token-metadata" aria-hidden="true">#</a> Get Token Metadata</h1><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/davecgh/go-spew/spew&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/common&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/program/metaplex/token_metadata&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// NFT in solana is a normal mint but only mint 1.</span>
	<span class="token comment">// If you want to get its metadata, you need to know where it stored.</span>
	<span class="token comment">// and you can use \`tokenmeta.GetTokenMetaPubkey\` to get the metadata account key</span>
	<span class="token comment">// here I take a random Degenerate Ape Academy as an example</span>
	mint <span class="token operator">:=</span> common<span class="token punctuation">.</span><span class="token function">PublicKeyFromString</span><span class="token punctuation">(</span><span class="token string">&quot;GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx&quot;</span><span class="token punctuation">)</span>
	metadataAccount<span class="token punctuation">,</span> err <span class="token operator">:=</span> token_metadata<span class="token punctuation">.</span><span class="token function">GetTokenMetaPubkey</span><span class="token punctuation">(</span>mint<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;faield to get metadata account, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// new a client</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>MainnetRPCEndpoint<span class="token punctuation">)</span>

	<span class="token comment">// get data which stored in metadataAccount</span>
	accountInfo<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetAccountInfo</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> metadataAccount<span class="token punctuation">.</span><span class="token function">ToBase58</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to get accountInfo, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// parse it</span>
	metadata<span class="token punctuation">,</span> err <span class="token operator">:=</span> token_metadata<span class="token punctuation">.</span><span class="token function">MetadataDeserialize</span><span class="token punctuation">(</span>accountInfo<span class="token punctuation">.</span>Data<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to parse metaAccount, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>metadata<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br><span class="line-number">32</span><br><span class="line-number">33</span><br><span class="line-number">34</span><br><span class="line-number">35</span><br><span class="line-number">36</span><br><span class="line-number">37</span><br><span class="line-number">38</span><br><span class="line-number">39</span><br><span class="line-number">40</span><br><span class="line-number">41</span><br></div></div>`,2);function p(e,o){return t}var l=n(s,[["render",p]]);export{l as default};
